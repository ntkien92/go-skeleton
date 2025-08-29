package model

import (
	"time"

	"github.com/google/uuid"
)

type Article struct {
	ModelUUID
	Title      string
	Content    string
	CategoryID uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time

	Category *Category `gorm:"foreignKey:CategoryID"`
}

type GetDetailArticleQueryParams struct {
	Id *string
}

func (p GetDetailArticleQueryParams) Map() (Article, error) {
	var artile Article
	if p.Id != nil {
		id, err := uuid.Parse(*p.Id)
		if err != nil {
			return artile, err
		}

		artile.ID = id
	}

	return artile, nil
}
