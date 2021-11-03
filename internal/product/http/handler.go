package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/internal/product"
)

type Handler struct {
	useCase product.UseCase
}

func NewHandler(us product.UseCase) *Handler {
	return &Handler{
		useCase: us,
	}
}

func (handler *Handler) GetHttp(ctx *gin.Context) {
	categorySlug := ctx.Param("categorySlug")
	if categorySlug == "" {
		ctx.HTML(http.StatusNotFound, "404.html", gin.H{
			"extra_message": "You should provide category, " +
				"if you want to lookup for products",
		})
		return
	}
	products, err := handler.useCase.GetProductsByCategorySlug(categorySlug)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "base/404.html", gin.H{
			"extra_message": err.Error(),
		})
		return
	}
	ctx.HTML(http.StatusOK, "product-list.html", gin.H{
		"object_list": products,
	})
}
