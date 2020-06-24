package app

import controller "github.com/farazmunj/bookstore_users-api/controllers"

func MapURLs() {
	router.GET("/ping", controller.Ping)
}
