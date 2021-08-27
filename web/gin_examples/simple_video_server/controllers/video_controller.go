package controllers

import (
	"gin_examples/simple_video_server/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

type VideoController interface {
	GetAll(context *gin.Context)

	Update(context *gin.Context)

	Create(context *gin.Context)

	Delete(context *gin.Context)
}

type controller struct {
	videos []models.Video
}

type generator struct {
	count int
	mtx   sync.Mutex
}

var g = &generator{}

func (g *generator) getNextId() int {
	g.mtx.Lock()
	g.count++
	g.mtx.Unlock()
	return g.count
}

func NewVideoController() VideoController {
	return &controller{videos: make([]models.Video, 0)}
}

func (c *controller) GetAll(context *gin.Context) {
	context.JSON(http.StatusOK, c.videos)
}

func (c *controller) Update(context *gin.Context) {
	var videoToUpdate models.Video
	if err := context.ShouldBindUri(&videoToUpdate); err != nil {
		context.String(http.StatusBadRequest, "bad request %v", err)
		return
	}
	if err := context.ShouldBind(&videoToUpdate); err != nil {
		context.String(http.StatusBadRequest, "bad request %v", err)
		return
	}
	for idx, video := range c.videos {
		if video.Id == videoToUpdate.Id {
			c.videos[idx] = videoToUpdate
			context.String(http.StatusOK, "success, video with id %d has been updated", videoToUpdate.Id)
			return
		}
	}
	context.String(http.StatusBadRequest, "bad request, cannot find video with %d to update", videoToUpdate.Id)
}

func (c *controller) Create(context *gin.Context) {
	video := models.Video{
		Id: g.getNextId(),
	}
	if err := context.BindJSON(&video); err != nil {
		context.String(http.StatusBadRequest, "bad request, %v", err)
		return
	}
	c.videos = append(c.videos, video)
	context.String(http.StatusOK, "success, new video id is %id", video.Id)
}

func (c *controller) Delete(context *gin.Context) {
	var videoToDelete models.Video
	if err := context.ShouldBindUri(&videoToDelete); err != nil {
		context.String(http.StatusBadRequest, "bad request, %v", err)
		return
	}
	for idx, video := range c.videos {
		if video.Id == videoToDelete.Id {
			c.videos = append(c.videos[0:idx], c.videos[idx:len(c.videos)-1]...)
			context.String(http.StatusOK, "success, video with id %d has been deleted", videoToDelete.Id)
			return
		}
	}
	context.String(http.StatusBadRequest, "bad request, cannot find videos with %d to delete,", videoToDelete.Id)
}
