package storage

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Book - We will be using this Book type to perform crud operations
type UserTwo struct {
	firstName string
	email     string
	salary    string
	lastName  string
}

func main() {

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	usersCollection := client.Database("testdb").Collection("users")
	// Insert One document
	user1 := UserTwo{"Some first name", "test@email.com", "10000", "Some last name"}
	insertResult, err := usersCollection.InsertOne(context.TODO(), user1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}
