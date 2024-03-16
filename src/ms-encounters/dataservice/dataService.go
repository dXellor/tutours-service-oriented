package dataservice

import (
	"tutours/soa/ms-encounters/model"

	"gorm.io/gorm"
)

type ICRUDRepository interface {
	GetAll() ([]model.Encounter, error)
	Get(id int) (model.Encounter, error)
	Create(l *model.Encounter) (model.Encounter, error)
	Delete(id int) error
	Update(l *model.Encounter) (model.Encounter, error)
}

type IEncounterRepository interface {
	ICRUDRepository
	Init(databaseConnection *gorm.DB)
}