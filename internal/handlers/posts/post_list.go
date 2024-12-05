package posts

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllPost(c *gin.Context){
	ctx:= c.Request.Context()
	pageIndexStr:= c.Query("pageIndex")
	pageSize:= c.Query("pageSize")

	pageIndex, err := strconv.Atoi(pageIndexStr)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error" :errors.New("invalid page index").Error(),
		})
		return
	}
	pagesize, err := strconv.Atoi(pageSize)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error" :errors.New("invalid page size").Error(),
		})
		return
	}
	response , err := h.postSVC.GetAllPost(ctx,pagesize,pageIndex)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error" :err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}