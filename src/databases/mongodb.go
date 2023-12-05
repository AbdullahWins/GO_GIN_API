// database/db.go
package database

import (
	"context"

	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client          *mongo.Client
	Database        *mongo.Database
	TodosCollection *mongo.Collection
	TasksCollection *mongo.Collection
)

func ConnectDatabase() error {

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	Client = client
	Database = client.Database(os.Getenv("DATABASE_NAME"))
	TodosCollection = Database.Collection("todosCollection")
	TasksCollection = Database.Collection("tasksCollection")

	return nil
}
