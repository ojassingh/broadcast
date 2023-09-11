package routes

import (
	"context"
	"fmt"
	"net/http"

	collection "github.com/TutorialEdge/realtime-chat-go-react/Collection"
	database "github.com/TutorialEdge/realtime-chat-go-react/databases"
	"github.com/TutorialEdge/realtime-chat-go-react/model"
	"github.com/gin-gonic/gin"
)

func CreateItem(c *gin.Context) {

	var newItem model.Item
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	var client = database.ConnectDB()
	var collection = collection.GetCollection(client)

	result, err := collection.InsertOne(context.TODO(), newItem)
	if err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"_id": result.InsertedID})

	database.DisconnectDB(client)
}
