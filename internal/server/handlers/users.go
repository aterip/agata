package handlers

import (
	"fmt"
	"net/http"
)

func (h *Handler) AuthUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Authentification")
}

func (h *Handler) RegistrationUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Registration")
}

func (h *Handler) EditUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Edit")
}

func (h *Handler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Delete")
}

func (h *Handler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get user")
}

func (h *Handler) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get users")
}
