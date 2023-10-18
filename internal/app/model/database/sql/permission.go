package sql

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Permission struct {
	// Migrated Fields
	ID        uint   `gorm:"primaryKey"`
	XID       string `gorm:"column:xid"`
	Name      string
	Slug      string
	Path      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Permission) BeforeCreate(tx *gorm.DB) error {
	if p.XID == "" {
		p.XID = uuid.NewString()
	}
	return nil
}

func (Permission) TableName() string {
	return "permissions"
}
