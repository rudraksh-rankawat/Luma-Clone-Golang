package main

import (
	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	err := server.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
