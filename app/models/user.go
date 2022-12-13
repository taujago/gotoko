package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            string `gorm:"size:36;not null, uniqueIndex; parimary_key"`
	Address       []Address
	FirstName     string `gorm:"size:100; not null"`
	LastName      string `gorm:"size:100; not nulll"`
	Email         string `gorm:"size:100; not null; uniqueIndex"`
	Password      string `gorm:"size:255; not null"`
	RememberToken string `gorm:"size:255; not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}
