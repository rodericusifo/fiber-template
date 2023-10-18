package sql

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	// Migrated Fields
	ID        uint   `gorm:"primaryKey"`
	XID       string `gorm:"column:xid"`
	Name      string
	Slug      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (r *Role) BeforeCreate(tx *gorm.DB) error {
	if r.XID == "" {
		r.XID = uuid.NewString()
	}
	return nil
}

func (Role) TableName() string {
	return "roles"
}
