package routes

import (
	"context"
	"net/http"

	collection "github.com/TutorialEdge/realtime-chat-go-react/Collection"
	database "github.com/TutorialEdge/realtime-chat-go-react/databases"
	"github.com/TutorialEdge/realtime-chat-go-react/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ReadItem(c *gin.Context) {

	var client = database.ConnectDB()
	var collection = collection.GetCollection(client)

	itemID := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(itemID)
	if err != nil {
		// Handle the error, e.g., invalid item ID format
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID format"})
		return
	}

	filter := bson.D{{"_id", objectID}}

	// cursor, err := collection.Find()

	var item model.Item
	err = collection.FindOne(context.TODO(), filter).Decode(&item)
	if err != nil {
		// Handle the error, e.g., item not found
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, item)

	database.DisconnectDB(client)
}

func ReadAll(c *gin.Context) {

	var client = database.ConnectDB()
	var collection = collection.GetCollection(client)

	filter := bson.D{}

	cursor, err := collection.Find(context.TODO(), filter)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer cursor.Close(context.TODO())

	var items []model.Item
	for cursor.Next(context.TODO()) {
		var item model.Item
		if err := cursor.Decode(&item); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		items = append(items, item)
	}

	c.JSON(http.StatusOK, items)

	database.DisconnectDB(client)
}
