// models/task.go
package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title  string             `json:"title" bson:"title"`
	Status string             `json:"status" bson:"status"`
}
