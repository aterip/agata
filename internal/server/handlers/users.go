package handlers

import (
	"fmt"
	"net/http"
)

func AuthUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Authentification")
}

func RegistrationUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Registration")
}

func EditUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Edit")
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Delete")
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get user")
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get users")
}
