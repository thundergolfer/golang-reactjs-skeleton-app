package datastores

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	"google.golang.org/api/iterator"

	"cloud.google.com/go/storage"
	"github.com/satori/go.uuid"
	"github.com/thundergolfer/12-factor/backend/types"
	"golang.org/x/net/context"
)

type GoogleCloudStorer struct {
	client     *storage.Client
	bucketName string
}

func NewGoogleCloudStorer(projectID string, bucketName string, ctx context.Context) *GoogleCloudStorer {
	storer := GoogleCloudStorer{
		bucketName: bucketName,
	}
	// Creates a client.
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return nil
	}

	storer.client = client

	return &storer
}

func (s *GoogleCloudStorer) ListTodos() types.Todos {
	ctx := context.Background()

	todos := types.Todos{}

	bucket := s.client.Bucket(s.bucketName)
	it := bucket.Objects(ctx, nil)
	for {
		objAttrs, err := it.Next()
		if err != nil && err != iterator.Done {
			log.Warn(err)
			continue
		}
		if err == iterator.Done {
			break
		}
		if objReader, err := bucket.Object(objAttrs.Name).NewReader(ctx); err != nil {
			log.Warn(err) // Just skip this todo
			continue
		} else {
			slurp, err := ioutil.ReadAll(objReader)
			objReader.Close()
			if err != nil {
				log.Warn(err) // Just skip
				continue
			}

			idStr := strings.TrimSuffix(objAttrs.Name, filepath.Ext(objAttrs.Name))
			if err != nil {
				log.Warn("Found non-numeric .txt filename in Todos bucket")
				continue // Just skip it
			}

			todos = append(todos, types.Todo{
				Text: string(slurp),
				Id:   idStr,
			})
		}

	}
	return todos
}

// FindTodo retrieves a Todo obj from Google Cloud Storage if it exists.
// If it doesn't, the method will return an empty Todo
func (s *GoogleCloudStorer) FindTodo(id string) types.Todo {
	ctx := context.Background()
	// Creates a Bucket instance.
	bucket := s.client.Bucket(s.bucketName)

	objReader, err := bucket.Object(id + ".txt").NewReader(ctx)
	if err == nil {
		log.Warn(err)
		return types.Todo{}
	}

	slurp, err := ioutil.ReadAll(objReader)
	objReader.Close()
	if err != nil {
		log.Warn(err)
		return types.Todo{}
	}

	return types.Todo{
		Text: string(slurp),
	}

}

// CreateTodo stores a new Todo in Google Cloud Storage
func (s *GoogleCloudStorer) CreateTodo(t types.Todo) types.Todo {
	ctx := context.Background()

	t.Id = uuid.NewV4().String()
	wc := s.client.Bucket(s.bucketName).Object(t.Id + ".txt").NewWriter(ctx)

	wc.ContentType = "text/plain"
	wc.ACL = []storage.ACLRule{{storage.AllUsers, storage.RoleReader}}
	if _, err := wc.Write([]byte(t.Text)); err != nil {
		// TODO: handle error.
		// Note that Write may return nil in some error situations,
		// so always check the error from Close.
		log.Warn(err)
		return types.Todo{} // TODO: should return an err
	}
	if err := wc.Close(); err != nil {
		// TODO: handle error.
		log.Warn(err)
		return types.Todo{} // TODO: should return an err
	}

	return t
}

// DestroyTodo removes a Todo from Google Cloud Storage
func (s *GoogleCloudStorer) DestroyTodo(id string) error {
	ctx := context.Background()
	bucket := s.client.Bucket(s.bucketName)
	todoFile := id + ".txt"
	if err := bucket.Object(todoFile).Delete(ctx); err != nil {
		// TODO: Handle error.
		return err
	}
	return nil
}
