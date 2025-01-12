package db

import (
	"context"
	"fmt"
	"log"
	"time"
	"workoutbot/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TeamsGetOne(teamName string) (*models.Team, error) {
	// Select the database and collection
	teams := GetCollection("workoutbot", "teams")

	// Define the filter for the document you want to find
	filter := bson.M{"team_name": teamName}

	// Context for query
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Query for a single document
	var team models.Team
	err := teams.FindOne(ctx, filter).Decode(&team)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No document found with the specified filter.")
			return nil, nil
		}
		// Return other errors
		return nil, fmt.Errorf("failed to find document: %w", err)
	} else {
		fmt.Println("Team found: ", team)
	}

	return &team, nil
}

func TeamsGetAll() ([]models.Team, error) {
	// Select the database and collection
	teams := GetCollection("workoutbot", "teams")

	// Define the filter for the document you want to find
	// Empty filter = Find all documents in the collection
	filter := bson.M{}

	// Context for query
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Define an empty slice to hold the results
	var results []models.Team

	// Query for a single document
	allTeams, err := teams.Find(ctx, filter)
	if err != nil {
		log.Printf("Failed to find documents: %v", err)
	}
	defer allTeams.Close(ctx)

	if err = allTeams.All(ctx, &results); err != nil {
		return nil, fmt.Errorf("failed to decode documents: %w", err)
	}

	// Print all documents
	fmt.Println("Documents in collection:")
	for _, doc := range results {
		fmt.Println(doc)
	}

	return results, nil
}

func TeamsSaveOne(teamName string) (primitive.ObjectID, error) {
	// Select the database and collection
	teams := GetCollection("workoutbot", "teams")

	// Context for query
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if team name already exists
	existingTeam, err := TeamsGetOne(teamName)
	if err != nil {
		log.Printf("Something happened when calling \"TeamsGetOne\": %v", err)
	}

	// We don't want duplicate teams. Return a message saying the team already exists.
	if !existingTeam.Id.IsZero() {
		return primitive.NilObjectID, fmt.Errorf("team name \"%v\" already exists. Try another team. Use !workoutbot teams to see the existing teams", teamName)
	}

	// Create a new team document
	team := models.Team{
		TeamName: teamName,
	}

	// Insert the document
	result, err := teams.InsertOne(ctx, team)
	if err != nil {
		log.Printf("Failed to add Team %v: %v", teamName, err)
	}

	// Print the ID of the inserted document
	insertedID := result.InsertedID.(primitive.ObjectID)
	fmt.Printf("Team \"%v\" added with ID: %s\n", teamName, insertedID.Hex())

	return insertedID, err
}
