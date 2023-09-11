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

	// Convert itemID to BSON ObjectId
	objectID, err := primitive.ObjectIDFromHex(itemID)
	if err != nil {
		// Handle the error, e.g., invalid item ID format
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID format"})
		return
	}

	var client = database.ConnectDB()
	var collection = collection.GetCollection(client)

	filter := bson.M{"_id": objectID}

	// Delete the item from the database
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		// Handle the error, e.g., item not found
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	if result.DeletedCount == 0 {
		// If DeletedCount is 0, the item was not found
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	// Item successfully deleted
	c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})

	database.DisconnectDB(client)
}
