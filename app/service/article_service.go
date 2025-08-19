package service

import "blog-api/interfaces"

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

func (s *ArticleService) GetList() error {
	return nil
}
