package routes

import (
	"database/sql"
	"net/http"

	"github.com/Roni6291/event_booking/models"
	"github.com/Roni6291/event_booking/utils"
	"github.com/gin-gonic/gin"
)

func signUp(db *sql.DB) gin.HandlerFunc {

	fn := func(context *gin.Context) {

		var user models.User
		err := context.ShouldBindJSON(&user)

		if err != nil {
			context.JSON(
				http.StatusBadRequest,
				gin.H{"message": "Couldn't parse request data"},
			)
			return
		}
		err = user.Save(db)
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
				"message": "User Created Successfully",
				"event":   user,
			},
		)
	}
	return gin.HandlerFunc(fn)

}

func login(db *sql.DB) gin.HandlerFunc {

	fn := func(context *gin.Context) {

		var user models.User
		err := context.ShouldBindJSON(&user)

		if err != nil {
			context.JSON(
				http.StatusBadRequest,
				gin.H{"message": "Couldn't parse request data"},
			)
			return
		}

		err = user.Validate(db)
		if err != nil {
			context.JSON(
				http.StatusUnauthorized,
				gin.H{"message": err.Error()},
			)
			return
		}
		key := "SuperSecretKey"
		token, err := utils.GenerateToken(user.Name, user.Id, key)
		if err != nil {
			context.JSON(
				http.StatusInternalServerError,
				gin.H{"message": "Couldn't Generate token"},
			)
			return
		}

		context.JSON(
			http.StatusOK,
			gin.H{
				"message": "Login Successful!!",
				"token":   token,
			},
		)
	}
	return gin.HandlerFunc(fn)
}
