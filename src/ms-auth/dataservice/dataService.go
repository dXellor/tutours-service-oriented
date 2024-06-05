package dataservice

import (
	"ms-auth/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Init(databaseConnection *gorm.DB)
	GetActiveByUsername(username string) (model.User, error)
	GetPerson(id int) (model.Person, error)
}
