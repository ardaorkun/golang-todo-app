package db

import (
	"context"
	"errors"
	"github.com/ardaorkun/go-todo-app/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func GetTasks() ([]models.Task, error) {
	ctx, cancel := getContext()
	defer cancel()

	cursor, err := tasksCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []models.Task
	for cursor.Next(ctx) {
		var task models.Task
		err := cursor.Decode(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func GetTask(id string) (*models.Task, error) {
	ctx, cancel := getContext()
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}

	var task models.Task
	err = tasksCollection.FindOne(ctx, filter).Decode(&task)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return &task, nil
}

func CreateTask(title string, description string) error {
	ctx, cancel := getContext()
	defer cancel()

	task := models.Task{
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	_, err := tasksCollection.InsertOne(ctx, task)
	if err != nil {
		return err
	}

	return nil
}

func UpdateTask(id string, title *string, description *string, completed *bool) error {
	ctx, cancel := getContext()
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}

	update := bson.M{"$set": bson.M{}}

	if title != nil {
		update["$set"].(bson.M)["title"] = *title
	}

	if description != nil {
		update["$set"].(bson.M)["description"] = *description
	}

	if completed != nil {
		update["$set"].(bson.M)["completed"] = *completed
	}

	_, err = tasksCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTask(id string) error {
	ctx, cancel := getContext()
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}

	_, err = tasksCollection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	return nil
}

func getContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}
