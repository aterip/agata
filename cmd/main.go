package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/aterip/agata/internal/config"
	"github.com/aterip/agata/internal/logging"
	"github.com/aterip/agata/internal/server"
	"github.com/go-chi/chi/v5"
)

func main() {
	logger := logging.GetLogger()

	logger.Info("Create router")
	router := server.RegisterHandlers(logger)

	logger.Info("Get server config")
	cfg := config.GetServerConfig()

	logger.Info("Start server")
	startServer(router, cfg)

	// dbActiveList := config.GetDataBasesList()
	// pgConfig := config.GetPostgresConfig()

}

func startServer(router chi.Router, cfg *config.ServerConfig) {

	var listener net.Listener
	var listenErr error
	log.Println("start server")

	if cfg.Listen.Type == "sock" {
		log.Printf("Server start on socket")
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))

		if err != nil {
			log.Fatal(err)
		}

		socketPath := path.Join(appDir, "app.sock")
		log.Printf("Listen socket")
		listener, listenErr = net.Listen("unix", socketPath)

	} else {

		log.Printf("Listen %s", cfg.Listen.Protocol)
		listener, listenErr = net.Listen(cfg.Listen.Protocol, fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))

	}
	if listenErr != nil {
		log.Fatal(listenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	server.Serve(listener)

}
