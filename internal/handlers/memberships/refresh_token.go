package memberships

import (
	"forumapp-restapi/internal/model/memberships"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Refresh(c *gin.Context){
	ctx := c.Request.Context()
	var request memberships.RefreshTokenRequest
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	userID:=c.GetInt64("userID")

	accesToken, err := h.membershipSVC.ValidateRefreshToken(ctx,userID,request)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			
		})
		return
	}
	c.JSON(http.StatusOK, memberships.RefreshResponse{
		AccesToken: accesToken,
	})
}