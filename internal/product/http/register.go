package http

import (
	"github.com/gin-gonic/gin"
	"shop/internal/product"
)

func RegisterHttpEndpoints(router *gin.RouterGroup, us product.UseCase) {
	handler := NewHandler(us)
	group := router.Group("/products")
	{
		group.GET("/:categorySlug", handler.GetHttp)
	}
}
