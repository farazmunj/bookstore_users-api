package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/farazmunj/bookstore_users-api/domain/users"
	"github.com/farazmunj/bookstore_users-api/services"
	"github.com/gin-gonic/gin"
)

//CreateUser function to create new suer
func CreateUser(c *gin.Context) {
	var user users.User

	bytes, err := ioutil.ReadAll(c.Request.Body)
	if nil != err {
		//TODO: handle error
		return
	}
	if err := json.Unmarshal(bytes, &user); nil != err {
		// TODO: unmarshal error
		fmt.Println(err.Error())
	}
	fmt.Printf("User: %+v", user)
	fmt.Println(err)
	fmt.Println(string(bytes))

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
