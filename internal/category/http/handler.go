package http

import (
	"net/http"
	"shop/dtos"
	"shop/internal/category"
	"shop/models"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase category.UseCase
}

func NewHandler(us category.UseCase) *Handler {
	return &Handler{
		useCase: us,
	}
}

func (handler *Handler) Get(ctx *gin.Context) {
	categories, err := handler.useCase.GetAll()
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			&dtos.Message{
				Body: err.Error(),
			},
		)
		return
	}
	ctx.JSON(
		http.StatusOK,
		toResponseArray(categories),
	)
}

func (handler *Handler) GetDetail(ctx *gin.Context) {
	categoryId := ctx.Param("id")
	if categoryId == "" {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			&dtos.Message{
				Body: "Provide id of required category",
			},
		)
		return
	}
	modelCategory, err := handler.useCase.GetCategory(categoryId)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			&dtos.Message{
				Body: err.Error(),
			},
		)
	}
	ctx.JSON(
		http.StatusOK,
		toResponse(modelCategory),
	)
}

func (handler *Handler) HttpGet(ctx *gin.Context) {
	categories, err := handler.useCase.GetAll()
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			&dtos.Message{
				Body: err.Error(),
			},
		)
		return
	}
	ctx.HTML(http.StatusOK, "category-list.html", gin.H{
		"object_list": categories,
	})
}

func toResponse(model *models.Category) *dtos.Category {
	return &dtos.Category{
		Id:          model.Id,
		Title:       model.Title,
		Slug:        model.Slug,
		Description: model.Description,
		ImageUrl:    model.Image.Url,
	}
}

func toResponseArray(models *[]models.Category) *[]dtos.Category {
	out := make([]dtos.Category, len(*models))
	for index, model := range *models {
		out[index] = *toResponse(&model)
	}
	return &out
}
