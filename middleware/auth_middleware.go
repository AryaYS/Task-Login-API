package middleware

import (
	"apilogin/model/response"
	"encoding/json"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware AuthMiddleware) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	if req.Header.Get("X-API-Key") == "RHS" {
		middleware.Handler.ServeHTTP(wr, req)
	} else {
		wr.Header().Set("Content-Type", "application/json")
		wr.WriteHeader(http.StatusUnauthorized)

		webResponse := response.WebResponse{
			Code:   http.StatusNotFound,
			Status: "UNAUTHORIZED",
		}
		wr.Header().Add("Content-Type", "application/json")
		encoder := json.NewEncoder(wr)
		err := encoder.Encode(webResponse)
		if err != nil {
			panic(err)
		}
	}
}
