package db

import (
	"context"
	"fmt"
	"log"
	"time"
	"workoutbot/internal/models"

	"go.mongodb.org/mongo-driver/bson"
)

func WorkoutCategoryGetAll() (map[string]models.WorkoutCategory, error) {
	nameToWorkoutCategoryMap := make(map[string]models.WorkoutCategory)
	// Select the database and collection
	workoutCategories := GetCollection("workoutbot", "workout_categories")

	// Define the filter for the document you want to find
	// Empty filter = Find all documents in the collection
	filter := bson.M{}

	// Context for query
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var results []models.WorkoutCategory

	// Query for All Categories
	allCategories, err := workoutCategories.Find(ctx, filter)
	if err != nil {
		log.Printf("Failed to find documents: %v", err)
	}
	defer allCategories.Close(ctx)

	if err = allCategories.All(ctx, &results); err != nil {
		return nil, fmt.Errorf("failed to decode documents: %w", err)
	}

	for _, category := range results {
		nameToWorkoutCategoryMap[category.CategoryName] = category
	}

	// return results, nil
	return nameToWorkoutCategoryMap, err
}
