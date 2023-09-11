package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://ojassingh:ojassingh@cluster0.dkatlvz.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)
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
