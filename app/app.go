package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lavinas-science/learn-users-api/utils/logger"
)

var (
	router = gin.Default()
)

func StartApllication() {
	mapURLs()

	logger.Info("About to Start APP")
	router.Run(":8080")

}
