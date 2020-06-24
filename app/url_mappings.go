package app

import (
	"github.com/farazmunj/bookstore_users-api/controllers/ping"
	"github.com/farazmunj/bookstore_users-api/controllers/users"
)

// MapURLs will add all routes
func MapURLs() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/search", users.SearchUser)
	//router.GET("/users/:user_id", controllers.GetUser)
	router.POST("/users", users.CreateUser)
}
