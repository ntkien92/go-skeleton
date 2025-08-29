package model

import "time"

type Category struct {
	ModelUUID
	Name      string
	Articles  []Article
	CreatedAt time.Time
	UpdatedAt time.Time
}
