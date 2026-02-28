package main

import (
	"log"
	"net/http"

	"github.com/aterip/agata/internal/server/router"
)

func main() {
	startServer()

}

func startServer() {

	router.RegisterHandlers()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
