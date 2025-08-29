package repository

import (
	"strings"

	"gorm.io/gorm"
)

const (
	Category = "Category"
)

func PreloadPath(parts ...string) string {
	return strings.Join(parts, ".")
}

func applyPreloads(db *gorm.DB, preloads []string) *gorm.DB {
	for _, p := range preloads {
		db = db.Preload(p)
	}

	return db
}
