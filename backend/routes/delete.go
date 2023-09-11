package routes

import (
	"context"
	"net/http"

	collection "github.com/TutorialEdge/realtime-chat-go-react/Collection"
	database "github.com/TutorialEdge/realtime-chat-go-react/databases"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	var client = database.ConnectDB()
	var collection = collection.GetCollection(client)

	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully", "result": result})

}
