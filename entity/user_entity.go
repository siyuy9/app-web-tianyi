package entity

// User : database model for users
// http://jinzhu.me/gorm/models.html#model-definition
// if you want to split null and "", you should use *string instead of string.
type User struct {
	commonFields
	Username     *string `gorm:"uniqueIndex;not null" json:"username"`
	Email        *string `gorm:"not null"`
	PasswordHash *string `gorm:"column:password;not null"`
	Bio          string  `gorm:"size:1024"`
	Image        *string
	// "has many" association
	// https://gorm.io/docs/has_many.html
	Roles []Role `gorm:"foreignKey:FollowingID"`
}

// Role : represent the user's role
type Role struct {
	commonFields
	Name        string
	Description string
	Permissions []Permission
}

// Permission : represent the user's permission
type Permission struct {
	commonFields
	Name         string
	Path         string
	Description  string
	Capabilities []Capability
}

// Capbability : represent the user's capability
type Capability struct {
	commonFields
	Name        string
	Description string
}
