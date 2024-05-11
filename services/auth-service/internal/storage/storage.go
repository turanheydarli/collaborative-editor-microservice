package storage

import (
	"context"

	domain "github.com/turanheydarli/collaborative-editor/services/auth-service/internal/models"
)

type Storage interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserByUsername(ctx context.Context, username string) (*domain.User, error)
}
