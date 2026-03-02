package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/aterip/agata/internal/server"
	"github.com/go-chi/chi/v5"
)

func main() {
	router := server.RegisterHandlers()
	startServer(router)

}

func startServer(router chi.Router) {

	log.Println("start server")

	listener, err := net.Listen("tcp", ":8081")

	if err != nil {
		log.Fatal(err)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	server.Serve(listener)

}
