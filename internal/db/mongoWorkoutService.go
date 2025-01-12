package db

import (
	"context"
	"fmt"
	"log"
	"time"
	"workoutbot/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func WorkoutsInsertMany(workouts []models.Workout) (*mongo.InsertManyResult, error) {
	// Select the database and collection
	workoutCollection := GetCollection("workoutbot", "workouts")

	// Context for query
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var workoutsInterface []interface{}
	for _, workout := range workouts {
		workoutsInterface = append(workoutsInterface, workout)
	}

	// Insert the documents
	results, err := workoutCollection.InsertMany(ctx, workoutsInterface)
	if err != nil {
		log.Printf("Failed for some reason %v", err)
	} else {
		// Print the ID of the inserted document
		for _, result := range results.InsertedIDs {
			fmt.Printf("Insert Workout with ID: %s\n", result.(primitive.ObjectID).Hex())
		}
	}

	return results, err
}
