package services

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func MongoDBInit(){
	mongoDbPassword := os.Getenv("MONGODB_PASSWORD")
	if mongoDbPassword == "" {
		log.Fatal("MONGODB_PASSWORD not set in environment")
	}

	mongoDbUserId := os.Getenv("MONGODB_USERID")
	if mongoDbUserId == "" {
		log.Fatal("MONGODB_USERID not set in environment")
	}

	//var ctx, cancel = context.WithTimeout(context.Background(), 20*time.Second)
	connectionString := fmt.Sprintf("mongodb+srv://%v:%v@workoutbot.ee1do.mongodb.net/?retryWrites=true&w=majority&appName=workoutbot", mongoDbUserId, mongoDbPassword) 

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	MongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Ping the database to ensure the connection is established
	if err = MongoClient.Ping(ctx, nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}
	log.Println("Connected to MongoDB")
  
	// Send a ping to confirm a successful connection
	if err := MongoClient.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
	  panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	databases, err := MongoClient.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
	 log.Fatal(err)
	}
	fmt.Println(databases)
  }

  // GetCollection returns a reference to a MongoDB collection
func GetCollection(database, collection string) *mongo.Collection {
	if MongoClient == nil {
		log.Fatal("MongoDB client is not initialized. Call Connect first.")
	}
	return MongoClient.Database(database).Collection(collection)
}  