package repository

import (
	"blog-api/model"
	"testing"

	"github.com/stretchr/testify/assert"
	_ "modernc.org/sqlite"
)

func TestSeed(t *testing.T) {
	gormDb := setupTestDB(t)
	gormDb.AutoMigrate(&model.Article{})

	db, err := gormDb.DB()
	assert.NoError(t, err)

	err = Seed(db)
	assert.NoError(t, err)

	testCases := []struct {
		name        string
		sqlString   string
		expectCount int
	}{
		{
			name:        "count articles",
			sqlString:   "SELECT COUNT(*) FROM articles",
			expectCount: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var count int
			err := db.QueryRow(tc.sqlString).Scan(&count)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectCount, count)
		})
	}
}
