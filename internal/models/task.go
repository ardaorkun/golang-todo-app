package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Completed   bool               `bson:"completed"`
	CreatedAt   time.Time          `bson:"createdAt"`
	CompletedAt time.Time          `bson:"completedAt"`
}
