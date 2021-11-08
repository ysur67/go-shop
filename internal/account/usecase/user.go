package usecase

import (
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"shop/internal/account"
	"shop/models"
	"time"
)

type AuthClaims struct {
	jwt.StandardClaims
	User *models.User `json:"user"`
}

type UseCase struct {
	repo           account.Repository
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

func NewAccountUseCase(
	accountRepo account.Repository,
	hashSalt string,
	signingKey []byte,
	tokenTtl time.Duration) *UseCase {
	return &UseCase{
		repo:           accountRepo,
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		expireDuration: time.Second * tokenTtl,
	}
}

func (useCase *UseCase) LoginUser(user models.User) (string, error) {
	user.Password = getEncodedPassword(user.Password, useCase.hashSalt)
	existingUser, err := useCase.repo.GetUser(models.User{Username: user.Username, Password: user.Password})
	if err != nil {
		return "", err
	}
	claims := AuthClaims{
		User: existingUser,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(useCase.expireDuration)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(useCase.signingKey)
}

func (useCase *UseCase) RegisterUser(user models.User) error {
	user.Password = getEncodedPassword(user.Password, useCase.hashSalt)
	return useCase.repo.CreateUser(user)
}

func (useCase *UseCase) ParseToken(accessToken string) (*models.User, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return useCase.signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.User, nil
	}

	return nil, err
}

func getEncodedPassword(pwd string, hash string) string {
	out := sha1.New()
	out.Write([]byte(pwd))
	out.Write([]byte(hash))
	return fmt.Sprintf("%x", out.Sum(nil))
}
