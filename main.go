package main

import (
	"JiaoNiBan-server/databases"
	"JiaoNiBan-server/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	databases.Init()
	defer databases.Close()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1:80"})
	r.GET("/api/latest", middlewares.Latest)

	webGroup := r.Group("/api/website")
	{
		webGroup.GET("/desc", middlewares.WebsiteDesc)
		webGroup.GET("/content", middlewares.WebsiteContent)
		webGroup.GET("/index", middlewares.WebsiteIndex)
	}

	r.NoRoute(func(c *gin.Context) {
		c.Status(404)
	})

	r.Run(":8080")
}
