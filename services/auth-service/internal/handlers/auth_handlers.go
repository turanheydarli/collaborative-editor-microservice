package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/turanheydarli/collaborative-editor/services/auth-service/internal/models"
	"github.com/turanheydarli/collaborative-editor/services/auth-service/internal/security"
	"github.com/turanheydarli/collaborative-editor/services/auth-service/internal/services"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var creds models.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, err := models.NewUser(creds.Username, creds.Email, creds.Password, []string{"User"})

	if err != nil {
		http.Error(w, "Unable to create user", http.StatusInternalServerError)
		return
	}

	if err = services.CreateUser(r.Context(), user); err != nil {
		http.Error(w, "Unable to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, err := services.GetUserByUsername(r.Context(), creds.Username)
	if err != nil || user == nil || !user.CheckPassowrd(creds.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := security.GenerateJWT(user.Username, user.Roles)

	if err != nil {
		fmt.Print(err.Error())
		http.Error(w, "Unable to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Login successful", "token": token})
}
