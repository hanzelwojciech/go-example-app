package users

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	ADMIN      = "ADMIN"
	SUPER_USER = "SUPER_USER"
	USER       = "USER"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Email    string    `gorm:"unique;not null;type:varchar(100)"`
	Password string    `gorm:"type:varchar(255);not null"`
	Role     string    `gorm:"type:user_role;default:'USER'"`
	IsActive bool      `gorm:"default:false"`
	gorm.Model
}

type UserInvitationToken struct {
	User   User
	UserID uuid.UUID `gorm:"type:uuid;not null"`
	Token  string    `gorm:"type:varchar(255);not null"`
}
