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

type IEncounterCompletionRepository interface {
	// ICRUDRepository
	Init(databaseConnection *gorm.DB)

	GetCompletedCountByUser(userId int) (int64, error)
	GetFailedCountByUser(userId int) (int64, error)
	GetCompletedCountByUserAndMonth(userId int, month int, year int) (int64, error)
	GetFailedCountByUserAndMonth(userId int, month int, year int) (int64, error)
}
