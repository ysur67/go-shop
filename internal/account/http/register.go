package http

import (
	"github.com/gin-gonic/gin"
	"shop/internal/account"
)

func RegisterApiEndpoints(router *gin.Engine, uc account.UseCase) {
	h := NewHandler(uc)

	authEndpoints := router.Group("/auth")
	{
		authEndpoints.POST("/register", h.RegisterUser)
		authEndpoints.POST("/login", h.LoginUser)
	}
}

func RegisterHttpEndpoints(router *gin.RouterGroup, uc account.UseCase) {
	handler := NewHandler(uc)
	httpEndpoints := router.Group("/")
	{
		httpEndpoints.GET("/profile", handler.AccountPage)
	}
}
