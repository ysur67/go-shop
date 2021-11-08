package http

import (
	"github.com/gin-gonic/gin"
	"shop/internal/account"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc account.UseCase) {
	h := NewHandler(uc)

	authEndpoints := router.Group("/auth")
	{
		authEndpoints.POST("/register", h.RegisterUser)
		authEndpoints.POST("/login", h.LoginUser)
	}
}
