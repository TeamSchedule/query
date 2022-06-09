package server

import (
	"github.com/gin-gonic/gin"
)

func SetupServer() *gin.Engine {
	r := gin.New()
	setLogger(r)
	setRoutes(r)
	return r
}
