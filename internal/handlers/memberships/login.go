package memberships

import (
	"forumapp-restapi/internal/model/memberships"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var request memberships.LoginRequest
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	accesToken, refreshToken,err := h.membershipSVC.Login(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	response := memberships.LoginResponse{
		AccesToken: accesToken,
		RefreshToken: refreshToken,

	}
	c.JSON(http.StatusOK, response)

}
