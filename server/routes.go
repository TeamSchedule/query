package server

import (
	"github.com/gin-gonic/gin"
	"query/server/handlers"
)

func setRoutes(router *gin.Engine) {
	v1 := router.Group("v1")
	{
		v1.GET("/team/:teamId", handlers.GetTeamById)
	}
}
