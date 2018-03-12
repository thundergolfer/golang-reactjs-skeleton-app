DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

source $DIR/.deployment_env_vars

cd $DIR/../frontend
yarn build

# need to run: go-bindata -prefix "frontend" -pkg main -o backend/bindata.go frontend/public/...
# to get frontend assets
cd $DIR/..
rm -f backend/bindata.go
go-bindata -prefix "frontend" -pkg main -o backend/bindata.go frontend/public/...

cd $DIR/../backend

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

cd $DIR/..
docker build -t gcr.io/${PROJECT_ID}/${IMAGE_NAME}:v${TAG} -f Dockerfile.twelvefactor .

# You can run the docker image locally with:
# `docker run -itp 8080:8080 thundergolfer/12-factor:latest`
