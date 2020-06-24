package users

import (
	"fmt"
	"net/http"

	"github.com/farazmunj/bookstore_users-api/domain/users"
	"github.com/farazmunj/bookstore_users-api/services"
	"github.com/gin-gonic/gin"
)

//CreateUser function to create new suer
func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); nil != err {
		fmt.Printf("User Json.Unmarshal error: %+v", err.Error())
		return
	}

	result, saveErr := services.CreateUser(user)
	if nil != saveErr {
		fmt.Printf("Save error %+v", saveErr.Error())
		return
	}

	c.JSON(http.StatusCreated, result)
}

//GetUser function to get a suer
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

// SearchUser wil find a user
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
