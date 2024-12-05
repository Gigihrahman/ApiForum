package posts

import (
	"errors"
	"forumapp-restapi/internal/model/posts"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UpsertUserActivity(c *gin.Context){
	ctx:=c.Request.Context()
	var request posts.UserActivityRequest
	if err:= c.ShouldBindBodyWithJSON(&request); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error" :err.Error(),
		})
		return
	}
	postIDstr:=c.Param("postID")
	postID, err :=strconv.ParseInt(postIDstr, 10, 64)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : errors.New(" postID pada param tidak valid").Error(),
		})
		return
	}
	userID :=c.GetInt64("userID")
	err = h.postSVC.UpsertUserActivity(ctx,postID,userID, request)
	if err!= nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}
	c.Status(http.StatusOK)


}