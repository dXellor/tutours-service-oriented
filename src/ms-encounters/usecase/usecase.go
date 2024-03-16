package usecase

import (
	"tutours/soa/ms-encounters/dataservice"
	"tutours/soa/ms-encounters/model"
)

type ICRUDService interface {
	//Init(crudRepository dataservice.ICRUDRepository)
	GetAll() ([]model.Encounter, error)
	Get(id int) (*model.Encounter, error)
	Create(l *model.Encounter) (*model.Encounter, error)
	Delete(id int) error
	Update(id int, l *model.Encounter) (*model.Encounter, error)
}

type IEncounterService interface {
	ICRUDService
	Init(crudRepository dataservice.IEncounterRepository)
}