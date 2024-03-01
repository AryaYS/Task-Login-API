package helper

import (
	"apilogin/exception"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func GetCookieUser(req *http.Request) *jwt.Token {
	cookie, err := req.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			panic(exception.NotFoundErrorF("Not Found any Cookie"))
		}
		panic(err)
	}

	tokenStr := cookie.Value
	tkn, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte("your_key"), nil
	})
	if err != nil {
		panic(err)
	}
	return tkn
}
