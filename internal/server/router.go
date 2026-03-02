package server

import (
	"github.com/aterip/agata/internal/server/handlers"
	"github.com/go-chi/chi/v5"
)

func RegisterHandlers() chi.Router {

	r := chi.NewRouter()
	r.Route("/user", func(r chi.Router) {
		r.Post("/", handlers.RegistrationUserHandler)
		r.Put("/{uid}", handlers.EditUserHandler)
		r.Delete("/{uid}", handlers.DeleteUserHandler)
		r.Get("/{uid}", handlers.GetUserHandler)

	})
	r.Get("/users", handlers.GetUsersHandler)
	r.Post("/auth/{uid}", handlers.AuthUserHandler)

	return r
}
