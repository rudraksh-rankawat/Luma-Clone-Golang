package routes

import (
	"example.com/rest-api/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func getEvent(context *gin.Context) {

	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fetch the events. Try again later!"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEventById(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fetch the events. Try again later!"})
		return
	}
	id := context.Param("id")

	for _, event := range events {
		eventId := fmt.Sprintf("%v", event.ID)
		if eventId == id {
			context.JSON(http.StatusOK, event)
			return
		}
	}
	context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to fetch the event by the ID. Try again later!"})
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "unable to parse body"})
		return
	}

	event.UserID = 1
	event.DateTime = time.Now()
	var id int64
	id, err = event.Save()
	event.ID = id
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to create the event. Try again later!"})
		return
	}
	context.JSON(http.StatusCreated, event)
}

func updateEvent(c *gin.Context) {
	//_, err := models.GetAllEvents()
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"message": "Unable to get events"})
	//	return
	//}

	//var id int64

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unable to parse event ID"})
		return
	}

	_, err = models.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event by id"})
		return
	}

	var updatedEvent models.Event
	err = c.ShouldBindJSON(&updatedEvent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot parse the body"})
		return
	}

	updatedEvent.ID = id
	err = updatedEvent.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cound not update event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "event updated successfully"})
}

func deleteEventById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "unable to parse event ID"})
	}

	var event models.Event
	event, err = models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event by id"})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"messgage": "could not delete event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Event of ID %v deleted successfully", id)})
}

func deleteAllEvents(ctx *gin.Context) {
	err := models.DeleteAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete all events. Try later!"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "successfully deleted all events!"})
}
