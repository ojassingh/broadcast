package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	secretKey := os.Getenv("MONGODB_URI")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(secretKey).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil
	}

	fmt.Println("Successfully connected to MongoDB")

	return client
}

func DisconnectDB(client *mongo.Client) error {
	if client == nil {
		return nil // No client to disconnect
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		return err
	}

	fmt.Println("MongoDB client disconnected successfully")
	return nil
}
