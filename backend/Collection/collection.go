package collection

import "go.mongodb.org/mongo-driver/mongo"

func GetCollection(client *mongo.Client) *mongo.Collection {
	// var client = database.ConnectDB()
	collection := client.Database("sample").Collection("sample")
	return collection
}
