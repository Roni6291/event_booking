package routes

import (
	"database/sql"
	"net/http"
	"strconv"

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

		token := context.Request.Header.Get("Authorization")
		if token == "" {
			context.JSON(
				http.StatusUnauthorized,
				gin.H{"message": "Not Authorized"},
			)
			return
		}

		var event models.Event
		err := context.ShouldBindJSON(&event)

		if err != nil {
			context.JSON(
				http.StatusBadRequest,
				gin.H{"message": "Couldn't parse request data"},
			)
			return
		}
		event.UserId = 1
		err = event.Save(db)
		if err != nil {
			context.JSON(
				http.StatusInternalServerError,
				gin.H{"message": "Couldn't save the user in DB"},
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

func updateEvent(db *sql.DB) gin.HandlerFunc {

	fn := func(context *gin.Context) {
		eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)
		if err != nil {
			context.JSON(
				http.StatusBadRequest,
				gin.H{"message": "Couldn't parse eventId"},
			)
			return
		}
		_, err = models.GetEventById(eventId, db)
		if err != nil {
			context.JSON(
				http.StatusInternalServerError,
				gin.H{"message": "eventId not found in db"},
			)
			return
		}

		var updatedEvent models.Event
		err = context.ShouldBindJSON(&updatedEvent)
		if err != nil {
			context.JSON(
				http.StatusBadRequest,
				gin.H{"message": err.Error()},
			)
			return
		}
		updatedEvent.Id = eventId
		err = updatedEvent.Update(db)
		if err != nil {
			context.JSON(
				http.StatusInternalServerError,
				gin.H{"message": "Couldn't update the event in DB"},
			)
			return
		}
		context.JSON(
			http.StatusOK,
			gin.H{
				"message": "Event Updated Successfully!",
				"event":   updatedEvent,
			},
		)

	}
	return gin.HandlerFunc(fn)
}

func deleteEvent(db *sql.DB) gin.HandlerFunc {

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

		err = event.Delete(db)
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
				"message": "Event Deleted Successfully!",
			},
		)

	}
	return gin.HandlerFunc(fn)
}
