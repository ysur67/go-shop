package http

import (
	"github.com/gin-gonic/gin"
	"shop/internal/category"
)

func RegisterApiEndpoints(router *gin.RouterGroup, us category.UseCase) {
	handler := NewHandler(us)
	group := router.Group("/category")
	{
		group.GET("", handler.Get)
		group.GET("/:id", handler.GetDetail)
	}
}

func RegisterHttpEndpoints(router *gin.RouterGroup, us category.UseCase) {
	handler := NewHandler(us)
	group := router.Group("/category")
	{
		group.GET("", handler.HttpGet)
	}
}
