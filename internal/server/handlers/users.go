package handlers

import (
	"fmt"
	"net/http"
)

func (h *handler) AuthUserHandler(w http.ResponseWriter, r *http.Request) {
	//h.logger.Info("User is authentification")
	fmt.Fprint(w, "Authentification")
}

func (h *handler) RegistrationUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Registration")
}

func (h *handler) EditUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Edit")
}

func (h *handler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Delete")
}

func (h *handler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get user")
}

func (h *handler) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get users")
}
