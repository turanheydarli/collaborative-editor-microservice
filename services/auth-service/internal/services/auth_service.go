package services

import (
	"context"

	"github.com/turanheydarli/collaborative-editor/services/auth-service/internal/models"
	"github.com/turanheydarli/collaborative-editor/services/auth-service/internal/storage"
	"github.com/turanheydarli/collaborative-editor/services/auth-service/internal/storage/mongodb"
)

var storageService storage.Storage

func InitializeStorage() {
	storageService = mongodb.NewMongoStorage("mongodb://localhost:27017", "collaborative_editor", "users")
}

func CreateUser(ctx context.Context, user *models.User) error {
	return storageService.CreateUser(ctx, user)
}

func GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	return storageService.GetUserByUsername(ctx, username)
}
