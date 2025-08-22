package repository

import (
	"blog-api/interfaces"
	"blog-api/model"
	"blog-api/perrors"
	"errors"

	"github.com/labstack/echo/v4"
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

func (r ArticleRepository) GetList(ctx echo.Context) ([]model.Article, []error) {
	var result []model.Article
	var errs []error

	query := r.db.Model(model.Article{}).Find(&result)
	if err := query.Error; err != nil {
		errs = append(errs, err)
		return nil, errs
	}

	return result, nil
}

func (r ArticleRepository) Create(ctx echo.Context, data model.Article) (string, []error) {
	var errs []error

	if err := r.db.Create(&data).Error; err != nil {
		errs = append(errs, err)
		return "", errs
	}

	return data.ID.String(), nil
}

func (r ArticleRepository) GetDetail(ctx echo.Context, params model.GetDetailArticleQueryParams) (*model.Article, []error) {
	var errs []error
	var result *model.Article

	if err := r.db.Where(params.Map()).First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errs = append(errs, perrors.New(
				perrors.NotFound,
				"artile not found",
				params,
			))
			return nil, errs
		}

		errs = append(errs, err)
		return nil, errs
	}

	return result, nil
}
