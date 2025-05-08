package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"golang_restful/model/domain"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sqlx.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sqlx.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sqlx.Tx, category domain.Category)
	FindById(ctx context.Context, tx *sqlx.Tx, categoryId int) (domain.Category, error)
	FindAll(ctx context.Context, tx *sqlx.Tx) []domain.Category
}
