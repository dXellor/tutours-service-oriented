package model

import userRole "ms-stakeholders/model/enum"

type User struct {
	Id                int `gorm:"primaryKey;autoIncrement"`
	Username          string
	Password          string
	Role              userRole.UserRole
	IsActive          bool
	Email             string
	IsBlocked         bool
	IsEnabled         bool
	VerificationToken string
}
