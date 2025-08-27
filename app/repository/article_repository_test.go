package repository

import (
	"blog-api/model"
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestGetList(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&model.Article{})

	db.Create(&model.Article{ModelUUID: model.ModelUUID{ID: uuid.New()}, Title: "Title", Content: "Content"})

	repo := NewArticleRepository(db)

	result, errs := repo.GetList(context.TODO())

	assert.Nil(t, errs)
	assert.Len(t, result, 1)
	assert.Equal(t, "Title", result[0].Title)
}
