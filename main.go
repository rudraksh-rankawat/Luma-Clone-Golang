package main

import (
	"fmt"
	"net/http"

	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	// CORS middleware
	server.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	routes.RegisterRoutes(server)

	if err := server.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}
