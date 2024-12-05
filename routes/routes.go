package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("events/:id", getEventById)
	server.GET("/events", getEvent)
	server.POST("/events", createEvent)
	server.PUT("events/:id", updateEvent)
	server.DELETE("events/:id", deleteEventById)
	server.DELETE("/events", deleteAllEvents)
}
