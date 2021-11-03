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

func (handler *Handler) GetProductsByCategorySlugHttp(ctx *gin.Context) {
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
		ctx.HTML(http.StatusInternalServerError, "404.html", gin.H{
			"extra_message": err.Error(),
		})
		return
	}
	ctx.HTML(http.StatusOK, "product-list.html", gin.H{
		"object_list": products,
	})
}

func (handler *Handler) GetProductDetailHttp(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.HTML(http.StatusNotFound, "404.html", gin.H{
			"extra_message": "You have to provide product id, if" +
				" you want to see products detail info",
		})
		return
	}
	product, err := handler.useCase.GetProductById(id)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "404.html", gin.H{
			"extra_message": err.Error(),
		})
		return
	}
	ctx.HTML(http.StatusOK, "product-detail.html", gin.H{
		"object": product,
	})
}
