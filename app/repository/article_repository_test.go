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

func setupTestDB(t *testing.T) *gorm.DB {
	t.Helper()

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	if err := db.AutoMigrate(&model.Article{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}

	return db
}

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

func TestCreate(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&model.Article{})

	repo := NewArticleRepository(db)
	result, errs := repo.Create(context.TODO(), model.Article{
		Title:   "Title",
		Content: "Content",
	})

	assert.Nil(t, errs)
	assert.NotEmpty(t, result)
}

func TestGetDetail(t *testing.T) {
	db := setupTestDB(t)
	repo := NewArticleRepository(db)

	// Insert valid record
	id, errs := repo.Create(context.TODO(), model.Article{
		Title:   "Title",
		Content: "Content",
	})
	assert.Nil(t, errs)

	notFoundId := uuid.New().String()
	invalidId := "not-found-id"

	testCases := []struct {
		name        string
		inputID     *string
		expectError bool
		expectNil   bool
		expectID    string
	}{
		{
			name:        "success found",
			inputID:     &id,
			expectError: false,
			expectNil:   false,
			expectID:    id,
		},
		{
			name:        "invalid id format",
			inputID:     &invalidId,
			expectError: true,
			expectNil:   true,
		},
		{
			name:        "not found",
			inputID:     &notFoundId,
			expectError: true,
			expectNil:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, errs := repo.GetDetail(context.TODO(), model.GetDetailArticleQueryParams{
				Id: tc.inputID,
			})

			if tc.expectError {
				assert.NotNil(t, errs)
			} else {
				assert.Nil(t, errs)
			}

			if tc.expectNil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, tc.expectID, result.ID.String())
			}
		})
	}
}
