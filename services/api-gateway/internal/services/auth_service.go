package services

import (
	"errors"
)

type User struct {
	Username string
	Password string
}

var mockUsers = []User{
	{"admin", "password123"},
}

func Authenticate(username, password string) (string, error) {
	for _, user := range mockUsers {
		if user.Username == username && user.Password == password {
			return "mocked-jwt-token", nil
		}
	}
	return "", errors.New("invalid credentials")
}
