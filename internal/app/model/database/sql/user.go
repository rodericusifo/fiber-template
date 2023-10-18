package sql

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	// Migrated Fields
	ID        uint   `gorm:"primaryKey"`
	XID       string `gorm:"column:xid"`
	Name      string
	Email     string
	Password  string
	RoleID    uint
	CreatedAt time.Time
	UpdatedAt time.Time

	// Relations
	Role Role
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.XID == "" {
		u.XID = uuid.NewString()
	}
	return nil
}

func (User) TableName() string {
	return "users"
}
