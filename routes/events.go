package routes

import (
	"net/http"
	"strconv"
	"strings"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func getEvent(c *gin.Context) {
	search := c.Query("search")
	date := c.Query("date")

	events, err := models.GetAllEvents(search, date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fetch events"})
		return
	}

	c.JSON(http.StatusOK, events)
}


func getEventById(context *gin.Context) {
	idParam := context.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching event"})
		return
	}
	if event.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	token := context.GetHeader("Authorization")
	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Missing authorization"})
		return
	}
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
	}

	userId, err := utils.VerifyJWT(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		return
	}

	var event models.Event
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}

	event.UserID = userId

	id, err := event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event"})
		return
	}
	event.ID = id

	context.JSON(http.StatusCreated, event)
}

func updateEvent(context *gin.Context) {
	idParam := context.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}
	if event.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	var updatedEvent models.Event
	if err := context.ShouldBindJSON(&updatedEvent); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}

	updatedEvent.ID = id
	if err := updatedEvent.Update(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

func deleteEventById(context *gin.Context) {
	idParam := context.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}
	if event.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	if err := event.Delete(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
