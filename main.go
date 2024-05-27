package main

import (
	"net/http"

	"github.com/Roni6291/event_booking/db"
	"github.com/Roni6291/event_booking/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		// do something
		context.JSON(
			http.StatusBadRequest,
			gin.H{"message": err.Error()},
		)
		return
	}
	event.Id = 1
	event.UserId = "roabrah"
	event.Save()

	context.JSON(
		http.StatusCreated,
		gin.H{
			"message": "Event Created Successfully",
			"event":   event,
		},
	)
}

func main() {
	db.InitDB("events.db", 5, 2)

	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":9090")
}
