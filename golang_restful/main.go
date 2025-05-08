package golang_restful

import (
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"golang_restful/app"
	"golang_restful/controller"
	"golang_restful/repository"
	"golang_restful/service"
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

}
