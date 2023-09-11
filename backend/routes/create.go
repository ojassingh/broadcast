package routes

import (
	"context"
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

	_, err := collection.InsertOne(context.TODO(), newItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newItem)

	database.DisconnectDB(client)
}
