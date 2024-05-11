package mongodb

import (
	"context"
	"log"
	"time"

	domain "github.com/turanheydarli/collaborative-editor/services/auth-service/internal/models"
	"github.com/turanheydarli/collaborative-editor/services/auth-service/internal/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStorage struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoStorage(uri, dbName, collectionName string) storage.Storage {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	return &MongoStorage{
		client:     client,
		collection: client.Database(dbName).Collection(collectionName),
	}
}

func (s *MongoStorage) CreateUser(ctx context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := s.collection.InsertOne(ctx, user)

	return err
}

func (s *MongoStorage) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var user domain.User

	err := s.collection.FindOne(ctx, map[string]interface{}{"username": username}).Decode(&user)

	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	return &user, err
}
