package service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"golang_restful/exception"
	"golang_restful/helper"
	"golang_restful/model/domain"
	"golang_restful/model/web"
	"golang_restful/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sqlx.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sqlx.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate}
}

func (service CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	errValidate := service.Validate.Struct(request)
	helper.PanicIfError(errValidate)

	tx, err := service.DB.Beginx()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{Name: request.Name}

	category = service.CategoryRepository.Create(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	errValidate := service.Validate.Struct(request)
	helper.PanicIfError(errValidate)

	tx, err := service.DB.Beginx()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	category.Name = request.Name

	category = service.CategoryRepository.Update(ctx, tx, category)
	return helper.ToCategoryResponse(category)
}

func (service CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Beginx()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CategoryRepository.Delete(ctx, tx, category)

}

func (service CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Beginx()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)

}

func (service CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Beginx()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryReponses(categories)
}
