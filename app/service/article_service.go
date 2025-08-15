package service

import "blog-api/interfaces"

type ArticleService struct{}

func NewArticleService() interfaces.ArticleServiceInterface {
	return &ArticleService{}
}

func (s *ArticleService) GetList() error {
	return nil
}
