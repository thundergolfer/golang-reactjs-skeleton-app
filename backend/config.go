package main

import (
	"os"
)

type config struct {
	datastoreType                string
	projectID                    string
	googleCloudStorageBucketName string
}

func newConfig() config {
	c := config{}

	c.datastoreType = getEnv("DATASTORE", "local")
	c.projectID = getEnv("PROJECT_ID", "twelve-factor-app")
	c.googleCloudStorageBucketName = getEnv("GOOGLE_CLOUD_BUCKET_NAME", "")

	return c
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
