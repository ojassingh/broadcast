package main

import (
	"fmt"
	"net/http"

	"github.com/TutorialEdge/realtime-chat-go-react/pkg/websocket"
	"github.com/TutorialEdge/realtime-chat-go-react/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// "github.com/gorilla/websocket"
	// "github.com/TutorialEdge/realtime-chat-go-react/routes"
)

func handleWebSocket(c *gin.Context) {

	w := c.Writer
	r := c.Request

	pool := websocket.NewPool()
	go pool.Start()

	fmt.Println("WebSocket Endpoint Hit")
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func setupRoutes() {

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // Replace with your frontend's URL
	r.Use(cors.New(config))

	r.GET("/ws", handleWebSocket)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	r.GET("/items", routes.ReadAll)
	r.GET("/items/:id", routes.ReadItem)
	r.POST("/items", routes.CreateItem)
	r.DELETE("/items/:id", routes.DeleteItem)

	if err := r.Run(":8000"); err != nil {
		fmt.Println("Gin server failed: ", err)
	}

}

func main() {

	setupRoutes()

	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("server failed: ", err)
	}

}
