package authinternal

import (
	"errors"
	"fmt"
	"gin-base/internal/constant"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("your_jwt_secret_key")

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func GenerateToken(user User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["name"] = user.Name
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyToken(tokenString string) (*User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return nil, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New(constant.ErrInvalidToken)
	}

	user := claims["userId"].(User) // Type assertion

	fmt.Println("[internal-auth]user: ", user)
	return &User{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}
