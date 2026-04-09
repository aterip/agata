package server

import (
	"github.com/aterip/agata/internal/server/handlers"
	"github.com/aterip/agata/internal/logging"
	"github.com/go-chi/chi/v5"
)

func RegisterHandlers(logger logging.Logger) chi.Router {
	var h := handlers.Newhandler(logger)
	r := chi.NewRouter()
	r.Route("/user", func(r chi.Router) {
		r.Post("/", h.RegistrationUserHandler)
		r.Put("/{uid}", h.EditUserHandler)
		r.Delete("/{uid}", h.DeleteUserHandler)
		r.Get("/{uid}", h.GetUserHandler)

	})
	r.Get("/users", h.GetUsersHandler)
	r.Post("/auth/{uid}", h.AuthUserHandler)

	return r
}
