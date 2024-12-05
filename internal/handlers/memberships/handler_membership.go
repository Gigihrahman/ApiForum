package memberships

import (
	"context"
	"forumapp-restapi/internal/middleware"
	"forumapp-restapi/internal/model/memberships"

	"github.com/gin-gonic/gin"
)

type membershipService interface{
	SignUp(ctx context.Context, req memberships.SignUpRequest) error
	Login(ctx context.Context, req memberships.LoginRequest) (string,string,error)
	ValidateRefreshToken(ctx context.Context,userID int64, request memberships.RefreshTokenRequest) (string, error)

}

type Handler struct {
	*gin.Engine
	membershipSVC membershipService
}

func NewHandler (api *gin.Engine, membershipSVC membershipService) *Handler{
	return &Handler{
		Engine: api,
		membershipSVC: membershipSVC,
	}
}
func(h *Handler) RegisterRoute(){
	route := h.Group("/membership")
	route.GET("/test", h.ping)
	route.POST("/user", h.SignUp)
	route.POST("/login", h.Login)
	routeRefresh:= h.Group("memberships")
	routeRefresh.Use(middleware.AuthRefreshMiddleware())
	routeRefresh.POST("/refresh", h.Refresh)
}