package routes

import (
	"context"
	"net/http"

	collection "github.com/TutorialEdge/realtime-chat-go-react/Collection"
	database "github.com/TutorialEdge/realtime-chat-go-react/databases"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteItem(c *gin.Context) {
	itemID := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(itemID)
	if err != nil {
		// Handle the error, e.g., invalid item ID format
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID format"})
		return
	}

	filter := bson.D{{"_id", objectID}}
	var client = database.ConnectDB()
	var collection = collection.GetCollection(client)

	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully", "result": result})

}
