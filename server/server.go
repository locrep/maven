package server

import (
	"github.com/gin-gonic/gin"
	"github.com/locrep/locrep-go/config"
)

func mavenMiddleware(c *gin.Context) {
	println(c.Request.URL.String())
	c.Next()
}

func NewServer(config config.Conf) *gin.Engine {
	gin.SetMode(config.Environment.BuildMode())

	server := gin.New()
	server.Use(mavenMiddleware)

	return server
}