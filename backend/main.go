package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	port := "8080"
	log.WithFields(log.Fields{
		"port": port,
	}).Info("Starting App")

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":"+port, router))
}
