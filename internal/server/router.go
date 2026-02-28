package server

import (
	"net/http"

	"github.com/aterip/agata/internal/server/handlers"
)

type Route struct {
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

func RegisterHandlers() {

	routes := []Route{
		{"POST", "/registery", handlers.RegistrationUserHandler},
		{"POST", "/auth", handlers.RegistrationUserHandler},
		{"GET", "/edituser", handlers.EditUserHandler},
		{"GET", "/registery", handlers.RegistrationUserHandler},
	}
}
