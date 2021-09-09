package app

import (
	"github.com/lavinas-science/learn-users-api/controllers/ping"
	"github.com/lavinas-science/learn-users-api/controllers/users"
)

func mapURLs() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	router.GET("/users/search", users.SearchUser)
	router.POST("/users", users.CreateUser)
}
