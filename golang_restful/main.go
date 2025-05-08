package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/jackc/pgx/v5/stdlib"
	"golang_restful/app"
	"golang_restful/controller"
	"golang_restful/helper"
	"golang_restful/middleware"
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

	router := app.NewRouter(categoryController)

	serve := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := serve.ListenAndServe()
	helper.PanicIfError(err)
}
