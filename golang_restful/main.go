package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/julienschmidt/httprouter"
	"golang_restful/app"
	"golang_restful/controller"
	"golang_restful/exception"
	"golang_restful/helper"
	"golang_restful/repository"
	"golang_restful/service"
	"net/http"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.POST("/api/categories", categoryController.Create)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	serve := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := serve.ListenAndServe()
	helper.PanicIfError(err)
}
