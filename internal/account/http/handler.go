package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/dtos"
	"shop/internal/account"
	"shop/models"
)

type Handler struct {
	useCase account.UseCase
}

func NewHandler(useCase account.UseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (handler *Handler) RegisterUser(ctx *gin.Context) {
	inp := new(dtos.RegisterParams)
	if err := ctx.BindJSON(inp); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err := handler.useCase.RegisterUser(models.User{
		Username: inp.Username,
		Password: inp.Password,
	})
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	ctx.JSON(http.StatusCreated, dtos.Message{Body: "Successfully created a new account"})
}

func (handler *Handler) LoginUser(ctx *gin.Context) {
	inp := new(dtos.LoginParams)
	if err := ctx.BindJSON(inp); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	token, err := handler.useCase.LoginUser(models.User{
		Username: inp.Username,
		Password: inp.Password,
	})
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
	ctx.JSON(http.StatusOK, dtos.Message{Body: token})
}
