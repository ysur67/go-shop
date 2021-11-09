package account

import (
	"errors"
	"shop/models"
)

var ErrorInvalidAccessToken = errors.New("invalid token")
var ContextUserKey = "user"

type UseCase interface {
	RegisterUser(user models.User) error
	LoginUser(user models.User) (string, error)
	ParseToken(accessToken string) (*models.User, error)
	GetUser(username string) (*models.User, error)
}
