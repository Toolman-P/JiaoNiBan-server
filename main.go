package main

import (
	"JiaoNiBan-server/databases"
	"JiaoNiBan-server/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	databases.Init()
	defer databases.Close()
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.GET("/latest", middlewares.Latest)

	webGroup := r.Group("/website")
	{
		webGroup.GET("/desc", middlewares.WebsiteDesc)
		webGroup.GET("/content", middlewares.WebsiteContent)
		webGroup.GET("/index", middlewares.WebsiteIndex)
	}

	pushGroup := r.Group("/push")
	{
		pushGroup.POST("/update", middlewares.PushUpdate)
	}

	r.Run()
}
