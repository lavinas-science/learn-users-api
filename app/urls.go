package app

import (
	"github.com/lavinas-science/learn-users-api/controllers/ping"
	"github.com/lavinas-science/learn-users-api/controllers/users"
)

func mapURLs() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.Get)
	router.GET("/users/search", users.Search)
	router.POST("/users", users.Create)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	router.GET("/internal/users/search", users.Search)
}
