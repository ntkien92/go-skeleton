package model

import (
	"time"

	"github.com/google/uuid"
)

type Article struct {
	ModelUUID
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GetDetailArticleQueryParams struct {
	Id *string
}

func (p GetDetailArticleQueryParams) Map() Article {
	var artile Article
	if p.Id != nil {
		id, _ := uuid.Parse(*p.Id)
		artile.ID = id
	}

	return artile
}
