package model

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ModelUUID struct {
	ID uuid.UUID `gorm:"primarykey,default:(-)"`
}

func (u *ModelUUID) BeforeCreate(_ *gorm.DB) (err error) {
	u.ID, err = uuid.NewV7()

	if err != nil {
		return errors.New("failed to generate uuid")
	}

	return
}
