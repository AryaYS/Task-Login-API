package helper

import (
	"apilogin/model/response"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(user response.Response_user) string {
	claims := jwt.MapClaims{
		"user_name": user.User_name,
		"role_id":   user.Role_id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte("your_key")
	tokenString, err := token.SignedString(secret)
	if err != nil {
		panic(err)
	}
	return tokenString
}
