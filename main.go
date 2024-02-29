package main

import (
	"apilogin/connection"
	"apilogin/controller"
	"apilogin/exception"
	"apilogin/repository"
	"apilogin/service"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := connection.ConnectDb()
	val := validator.New()
	repo := repository.NewRepo()
	srv := service.NewService(repo, db, val)
	ctrl := controller.NewController(srv)
	router := httprouter.New()

	router.POST("/api/regis", ctrl.RegisterController)
	router.POST("/api/login", ctrl.LoginController)
	router.GET("/api/logout", ctrl.LogOut)
	router.GET("/api/GetByRole/:id", ctrl.GetAllWorkerBasedOnRole)

	router.PanicHandler = exception.ErrorHandler
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
