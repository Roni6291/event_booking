package main

import (
	"github.com/Roni6291/event_booking/db"
	"github.com/Roni6291/event_booking/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	DB := db.InitDB("events.db", 5, 2)

	router := gin.Default()
	routes.RegisterRoutes(router, DB)
	router.Run(":9090")
}
