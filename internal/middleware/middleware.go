package middleware

import (
	"errors"
	"forumapp-restapi/internal/configs"
	"forumapp-restapi/pkg/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc{
	secretKey := configs.Get().Service.SecretJWT
	return func(c *gin.Context){
		header:= c.Request.Header.Get("Authorization")
		header= strings.TrimSpace(header)
		if header == ""{
			c.AbortWithError(http.StatusUnauthorized, errors.New("missing Token"))
			return
		}

		userID,username, err:=jwt.ValidateToken(header,secretKey)
		

		if err!=nil{
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		c.Set("userID",userID)
		c.Set("username",username)
		c.Next()

	}
}


func AuthRefreshMiddleware() gin.HandlerFunc{
	secretKey := configs.Get().Service.SecretJWT
	return func(c *gin.Context){
		header:= c.Request.Header.Get("Authorization")
		header= strings.TrimSpace(header)
		if header == ""{
			c.AbortWithError(http.StatusUnauthorized, errors.New("missing Token"))
			return
		}

		userID,username, err:=jwt.ValidateTokenWithoutExpiry(header,secretKey)
		

		if err!=nil{
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		c.Set("userID",userID)
		c.Set("username",username)
		c.Next()

	}
}