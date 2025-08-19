package repository

import (
	"blog-api/interfaces"
	"blog-api/model"
	"context"

	"gorm.io/gorm"
)

type ArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(
	db *gorm.DB,
) interfaces.ArticleRepositoryInterface {
	return &ArticleRepository{
		db: db,
	}
}

func (r ArticleRepository) GetList(ctx context.Context) ([]model.Article, error) {
	return nil, nil
}
