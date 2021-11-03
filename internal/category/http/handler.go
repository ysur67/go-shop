package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/internal/category"
	"shop/models"
)

type Message struct {
	Body string `json:"body"`
}

type Category struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
}

type Handler struct {
	useCase category.UseCase
}

func NewHandler(us category.UseCase) *Handler {
	return &Handler{
		useCase: us,
	}
}

func (handler *Handler) GetDetail(ctx *gin.Context) {
	categoryId := ctx.Param("id")
	if categoryId == "" {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			&Message{
				"Provide id of required category",
			},
		)
		return
	}
	modelCategory, err := handler.useCase.GetCategory(ctx, categoryId)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			&Message{
				err.Error(),
			},
		)
	}
	ctx.JSON(
		http.StatusOK,
		toResponse(modelCategory),
	)
}

func toResponse(model *models.Category) *Category {
	return &Category{
		Id:          model.Id,
		Title:       model.Title,
		Slug:        model.Slug,
		Description: model.Description,
		ImageUrl:    model.Image.Url,
	}
}
