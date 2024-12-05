package posts

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetPostByID(c *gin.Context){
	ctx:= c.Request.Context()
	postIDstr:= c.Param("postID")
	postID, err := strconv.ParseInt(postIDstr,10,64)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : errors.New("invalid post Id").Error(),
		})
		return
	}
	response, err := h.postSVC.GetPostByID(ctx,postID)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}