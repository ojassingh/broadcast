package routes

import (
	"context"
	"net/http"

	collection "github.com/TutorialEdge/realtime-chat-go-react/Collection"
	database "github.com/TutorialEdge/realtime-chat-go-react/databases"
	"github.com/TutorialEdge/realtime-chat-go-react/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func ReadItem(c *gin.Context) {

	var client = database.ConnectDB()
	var collection = collection.GetCollection(client)

	filter := bson.D{}

	// cursor, err := collection.Find()

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.TODO())

	// var items []model.Item
	// for cursor.Next(context.TODO()) {
	// 	var item model.Item
	// 	if err := cursor.Decode(&item); err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	items = append(items, item)
	// }

	// c.JSON(http.StatusOK, items)

	c.JSON(200, "Hi")

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
