package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB) {
	//  GET routes
	router.GET("/events", getEvents(db))
	router.GET("/events/:eventId", getEvent(db))
	// POST routes
	router.POST("/events", createEvent(db))
	router.POST("/signup", signUp(db))
	router.POST("/login", login(db))
	// PUT routes
	router.PUT("/events/:eventId", updateEvent(db))
	// DELETE routes
	router.DELETE("/events/:eventId", deleteEvent(db))
}
