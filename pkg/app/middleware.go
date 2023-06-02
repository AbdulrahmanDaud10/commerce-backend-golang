package app

import (
	"errors"
	"fmt"
	"os"

	"github.com/anthdm/weavebox"
	"github.com/golang-jwt/jwt"
)

var ErrUnAuthenticated = errors.New("unAuthorized")

type AdminAuthMiddleware struct{}

func (mw *AdminAuthMiddleware) Authenticate(ctx *weavebox.Context) error {
	tokenString := ctx.Header("x-api-token")
	if len(tokenString) == 0 {
		return ErrUnAuthenticated
	}
	token, err := ParseJWT(tokenString)
	if err != nil {
		return ErrUnAuthenticated
	}
	if !token.Valid {
		return ErrUnAuthenticated
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return ErrUnAuthenticated
	}

	fmt.Println(claims)
	fmt.Println("guarding the admin routes")
	return nil
}

func ParseJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
}
