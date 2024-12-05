package posts

import (
	"forumapp-restapi/internal/model/posts"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePost(c *gin.Context){
	ctx := c.Request.Context()

	var request posts.CreatePostRequest
	if err:= c.ShouldBindBodyWithJSON(&request); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error" :err.Error(),
		})
		return
	}
	userID:= c.GetInt64("userID")
	err := h.postSVC.CreatePost(ctx,userID,request)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return
	}
	c.Status(http.StatusCreated)
}