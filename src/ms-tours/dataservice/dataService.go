package dataservice

import (
	"gorm.io/gorm"
)

type ICRUDRepository interface {
}

type ITourRepository interface {
	ICRUDRepository
	Init(databaseConnection *gorm.DB)
}
