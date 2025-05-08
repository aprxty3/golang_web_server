package repository

import (
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
	"golang_restful/helper"
	"golang_restful/model/domain"
)

type CategoryRepositoryImpl struct{}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Create(ctx context.Context, tx *sqlx.Tx, category domain.Category) domain.Category {
	SQL := "insert into category(name) values ($1) returning id"
	var id int
	err := tx.QueryRowContext(ctx, SQL, category.Name).Scan(&id)
	helper.PanicIfError(err)

	category.Id = id
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sqlx.Tx, category domain.Category) domain.Category {
	SQL := "update category set name = $1 where id = $2"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sqlx.Tx, category domain.Category) {
	SQL := "delete from category where id = $1"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sqlx.Tx, categoryId int) (domain.Category, error) {
	SQL := "select id, name from category where id = $1"
	row := tx.QueryRowContext(ctx, SQL, categoryId)

	category := domain.Category{}
	err := row.Scan(&category.Id, &category.Name)
	if err != nil {
		return category, errors.New("category is not found")
	}

	return category, nil
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sqlx.Tx) []domain.Category {
	SQL := "select id, name from category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}

	return categories
}
