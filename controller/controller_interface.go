package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	LoginController(wr http.ResponseWriter, req *http.Request, param httprouter.Params)
	RegisterController(wr http.ResponseWriter, req *http.Request, param httprouter.Params)
	GetAllWorkerBasedOnRole(wr http.ResponseWriter, r *http.Request, params httprouter.Params)
	LogOut(wr http.ResponseWriter, req *http.Request, params httprouter.Params)
}
