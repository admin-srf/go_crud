package services

import (
	"fmt"
	"time"

	"github.com/admin-srf/go_crud/schemas"
	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	Id    string `json:"id"`
	First string `json:"first"`
	Last  string `json:"last"`
	jwt.StandardClaims
}

func NewAccessToken(user schemas.User) (string, error) {
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()
	tk := &schemas.Token{
		UserID: user.ID,
		Name:   user.Name,
		Email:  user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, err := token.SignedString([]byte("secret"))

	return tokenString, err
}

func ParseToken(tokenString string) (*schemas.Token, error) {

	token, err := jwt.ParseWithClaims(tokenString, &schemas.Token{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})

	if claims, ok := token.Claims.(*schemas.Token); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}

}
