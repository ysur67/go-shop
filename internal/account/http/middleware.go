package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/internal/account"
	"strings"
)

type AuthMiddleware struct {
	useCase account.UseCase
}

func NewAuthMiddleware(useCase account.UseCase) gin.HandlerFunc {
	return (&AuthMiddleware{
		useCase: useCase,
	}).Handle
}

func (m *AuthMiddleware) Handle(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if headerParts[0] != "Bearer" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, err := m.useCase.ParseToken(headerParts[1])
	if err != nil {
		status := http.StatusInternalServerError
		if err == account.ErrorInvalidAccessToken {
			status = http.StatusUnauthorized
		}

		c.AbortWithStatus(status)
		return
	}

	c.Set(account.ContextUserKey, user)
}
