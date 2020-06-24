package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//CreateUser function to create new suer
func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

//GetUser function to get a suer
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

// SearchUser wil find a user
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
