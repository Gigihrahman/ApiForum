package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ping(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"message" :" testing"})
}