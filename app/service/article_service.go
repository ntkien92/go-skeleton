package service

import (
	"blog-api/dto"
	"blog-api/interfaces"
	"blog-api/model"
	"errors"

	"github.com/labstack/echo/v4"
)

type ArticleService struct {
	articleRepository interfaces.ArticleRepositoryInterface
}

func NewArticleService(
	articleRepository interfaces.ArticleRepositoryInterface,
) interfaces.ArticleServiceInterface {
	return &ArticleService{
		articleRepository: articleRepository,
	}
}

func (s *ArticleService) GetList(c echo.Context) ([]dto.ArticleResponse, []error) {
	articles, errs := s.articleRepository.GetList(c)
	if errs != nil {
		return nil, errs
	}

	response := dto.NewListArticleResponse(articles)
	return response, nil
}

func (s *ArticleService) GetDetail(ctx echo.Context, params dto.GetArticleDetailRequest) (*dto.ArticleResponse, []error) {
	article, errs := s.articleRepository.GetDetail(ctx, model.GetDetailArticleQueryParams{
		Id: &params.Id,
	})

	if errs != nil {
		return nil, errs
	}

	response := dto.NewArticleDetailResponse(*article)

	return &response, nil
}

func (s *ArticleService) Create(ctx echo.Context, data dto.CreateArticleRequest) (*dto.ArticleResponse, []error) {
	id, errs := s.articleRepository.Create(ctx, model.Article{
		Title:   data.Title,
		Content: data.Content,
	})

	if errs != nil {
		return nil, errs
	}

	if id == "" {
		errs = append(errs, errors.New("Can't get id"))
	}

	article, errs := s.articleRepository.GetDetail(ctx, model.GetDetailArticleQueryParams{
		Id: &id,
	})

	if errs != nil {
		return nil, errs
	}

	response := dto.NewArticleDetailResponse(*article)

	return &response, nil
}
