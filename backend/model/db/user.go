package db

import (
	"gorm.io/gorm"
)

// http://jinzhu.me/gorm/models.html#model-definition
// if you want to split null and "", you should use *string instead of string.
type UserModel struct {
	gorm.Model
	Username     string `gorm:"uniqueIndex;not null"`
	Email        string `gorm:"not null"`
	PasswordHash string `gorm:"column:password;not null"`
	Bio          string `gorm:"size:1024"`
	Image        *string
	// "has many" association
	// https://gorm.io/docs/has_many.html
	Roles []RoleModel `gorm:"foreignKey:FollowingID"`
}
type RoleModel struct {
	gorm.Model
	Name        string
	Description string
	Permissions []PermissionModel
}

type PermissionModel struct {
	gorm.Model
	Name         string
	Path         string
	Description  string
	Capabilities []CapabilityModel
}

type CapabilityModel struct {
	gorm.Model
	Name        string
	Description string
}
