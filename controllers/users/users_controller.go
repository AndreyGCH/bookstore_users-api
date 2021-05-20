package users

import (
	"github.com/gin-gonic/gin"
	"golang_api/bookstore_users-api/domain/users"
	"golang_api/bookstore_users-api/services"
	"golang_api/bookstore_users-api/utils/errors"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		//TODO: handle user creation error
		return
	}
	c.JSON(http.StatusNotImplemented, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}

//func SearchUser(c *gin.Context) {
//	c.String(http.StatusNotImplemented, "implement me")
//}
