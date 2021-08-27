package main

import (
	"gin_examples/simple_video_server/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/ping", func(context *gin.Context) {
		context.String(200, "%s", "pong")
	})

	server.StaticFile("/video", "resources/1617717463072884.mp4")

	videoController := controllers.NewVideoController()

	videoGroup := server.Group("/videos")
	{
		videoGroup.GET("/", videoController.GetAll)
		videoGroup.PUT("/:id", videoController.Update)
		videoGroup.POST("/", videoController.Create)
		videoGroup.DELETE("/:id", videoController.Delete)
	}

	server.Run(":12000")
}
