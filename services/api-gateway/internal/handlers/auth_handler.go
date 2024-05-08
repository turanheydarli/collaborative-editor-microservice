package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const authServiceURL = "http://localhost:8086"

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	loginURL := authServiceURL + "/login"
	reqBody, _ := json.Marshal(creds)

	resp, err := http.Post(loginURL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil || resp.StatusCode != http.StatusOK {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
}
