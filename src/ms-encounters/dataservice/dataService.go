package dataservice

import (
	"tutours/soa/ms-encounters/model"

	"tutours/soa/ms-encounters/model/enum"

	"gorm.io/gorm"
)

type ICRUDRepository interface {
	GetAll() ([]model.Encounter, error)
	Get(id int) (model.Encounter, error)
	Create(encounter *model.Encounter) (model.Encounter, error)
	Delete(id int) error
	Update(encounter *model.Encounter) (model.Encounter, error)
}

type IEncounterRepository interface {
	ICRUDRepository
	Init(databaseConnection *gorm.DB)
	GetApprovedByStatus(status enum.EncounterStatus) ([]model.Encounter, error) 
	GetByUser(userId int) ([]model.Encounter, error)
	GetTouristCreatedEncounters() ([]model.Encounter, error)
}