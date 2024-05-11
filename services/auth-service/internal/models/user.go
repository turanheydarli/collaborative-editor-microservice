package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id                 primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username           string             `json:"username" bson:"username"`
	Email              string             `json:"email" bson:"email"`
	PasswordHash       string             `json:"-" bson:"passwordHash"`
	Roles              []string           `json:"roles" bson:"roles"`
	CreatedAt          time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt          time.Time          `json:"updatedAt" bson:"updatedAt"`
	PasswordResetToken string             `json:"-" bson:"passwordResetToken"`
	EmailVerified      bool               `json:"emailVerified" bson:"emailVerified"`
}

func NewUser(username, email, password string, roles []string) (*User, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	return &User{
		Username:      username,
		Email:         email,
		PasswordHash:  string(hashedPassword),
		Roles:         roles,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		EmailVerified: false,
	}, nil
}

func (u *User) CheckPassowrd(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))

	return err == nil
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
