package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func MongoGetOneTeam(teamName string){
	log.Printf("Setting DB and Collection")
	// Select the database and collection
	// db := MongoClient.Database("workoutbot")
	// teams := db.Collection("teams")

	teams := GetCollection("workoutbot", "teams")

	log.Printf("Defining Filter")
	// Define the filter for the document you want to find
	filter := bson.M{"team_name": teamName}

	log.Printf("Defining Context for query")
	// Context for query
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Printf("Starting Single Query")
	// Query for a single document
	var team bson.M
	err := teams.FindOne(ctx, filter).Decode(&team)
	if err != nil{
		if err == mongo.ErrNoDocuments {
			fmt.Println("No document found with the specified filter.")
		} else {
			log.Fatalf("Failed to query single document: %v", err)
		}
	} else {
		fmt.Println("Team found: ", team)
	}
  }