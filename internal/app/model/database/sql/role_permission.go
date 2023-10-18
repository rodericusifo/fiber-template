package sql

import (
	"time"
)

type RolePermission struct {
	// Migrated Fields
	ID           uint `gorm:"primaryKey"`
	RoleID       uint
	PermissionID uint
	CreatedAt    time.Time
	UpdatedAt    time.Time

	// Relations
	Role       Role       `gorm:"constraint:OnDelete:CASCADE;"`
	Permission Permission `gorm:"constraint:OnDelete:CASCADE;"`
}

func (RolePermission) TableName() string {
	return "role_permissions"
}
