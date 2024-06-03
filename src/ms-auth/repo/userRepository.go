package repo

import (
	"ms-stakeholders/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	DatabaseConnection *gorm.DB
}

func (userRepository *UserRepository) Init(databaseConnection *gorm.DB) {
	userRepository.DatabaseConnection = databaseConnection
}

func (repo *UserRepository) GetActiveByUsername(username string) (model.User, error) {
	var user model.User
	dbResult := repo.DatabaseConnection.Where("username = ?", username).Find(&user)
	if dbResult.Error != nil {
		return user, dbResult.Error
	}
	return user, nil
}
func (repo *UserRepository) GetPerson(id int) (model.Person, error) {
	var person model.Person
	dbResult := repo.DatabaseConnection.Where("user_id = ?", id).Find(&person)
	if dbResult.Error != nil {
		return person, dbResult.Error
	}
	return person, nil
}
