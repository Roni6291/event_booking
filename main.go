package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/Roni6291/event_booking/db"
	"github.com/Roni6291/event_booking/models"
	"github.com/gin-gonic/gin"
)

func getEvent(db *sql.DB) gin.HandlerFunc {

	fn := func(context *gin.Context) {
		eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)
		if err != nil {
			context.JSON(
				http.StatusBadRequest,
				gin.H{"message": "Couldn't parse eventId"},
			)
			return
		}
		event, err := models.GetEventById(eventId, db)
		if err != nil {
			context.JSON(
				http.StatusInternalServerError,
				gin.H{"message": "eventId not found in db"},
			)
			return
		}
		context.JSON(
			http.StatusOK,
			gin.H{
				"message": "Event Fetched Successfully!",
				"event":   event,
			},
		)
	}
	return gin.HandlerFunc(fn)
}

func getEvents(db *sql.DB) gin.HandlerFunc {

	fn := func(context *gin.Context) {
		events, err := models.GetAllEvents(db)
		if err != nil {
			context.JSON(
				http.StatusInternalServerError,
				gin.H{"message": err.Error()},
			)
			return
		}
		context.JSON(http.StatusOK, events)
	}
	return gin.HandlerFunc(fn)

}

func createEvent(db *sql.DB) gin.HandlerFunc {

	fn := func(context *gin.Context) {
		var event models.Event
		err := context.ShouldBindJSON(&event)

		if err != nil {
			context.JSON(
				http.StatusBadRequest,
				gin.H{"message": err.Error()},
			)
			return
		}
		event.UserId = "roabrah"
		err = event.Save(db)
		if err != nil {
			context.JSON(
				http.StatusInternalServerError,
				gin.H{"message": err.Error()},
			)
			return
		}

		context.JSON(
			http.StatusCreated,
			gin.H{
				"message": "Event Created Successfully",
				"event":   event,
			},
		)
	}
	return gin.HandlerFunc(fn)

}

func main() {
	DB := db.InitDB("events.db", 5, 2)

	server := gin.Default()

	server.GET("/events", getEvents(DB))
	server.GET("/events/:eventId", getEvent(DB))
	server.POST("/events", createEvent(DB))
	server.Run(":9090")
}
