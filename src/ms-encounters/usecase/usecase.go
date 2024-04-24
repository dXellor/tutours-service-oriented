package usecase

import (
	"tutours/soa/ms-encounters/dataservice"
	"tutours/soa/ms-encounters/model"
	"tutours/soa/ms-encounters/model/dto"
	"tutours/soa/ms-encounters/model/enum"
)

type ICRUDService interface {
	//Init(crudRepository dataservice.ICRUDRepository)
	GetAll() ([]model.Encounter, error)
	Get(id int) (*model.Encounter, error)
	Create(encounter *model.Encounter) (*model.Encounter, error)
	Delete(id int) error
	Update(id int, encounter *model.Encounter) (*model.Encounter, error)
}

type IEncounterService interface {
	ICRUDService
	Init(crudRepository dataservice.IEncounterRepository)
	GetApprovedByStatus(status enum.EncounterStatus) ([]model.Encounter, error)
	GetByUser(userId int) ([]model.Encounter, error)
	GetTouristCreatedEncounters() ([]model.Encounter, error)
	Approve(encounter *model.Encounter) (*model.Encounter, error)
	Decline(encounter *model.Encounter) (*model.Encounter, error)
}

type IEncounterStatsService interface {
	// ICRUDService
	Init(crudRepository dataservice.IEncounterCompletionRepository)
	GetEncounterStatsByUser(userId int) (*dto.EncounterStats, error)
	GetEncounterYearStatsByUser(userId int, year int) (*dto.EncounterYearStats, error)
}

type IEncounterCompletionService interface {
	Init(crudRepository dataservice.IEncounterCompletionRepository)
	GetByUser(userId int) ([]model.EncounterCompletion, error)
	StartEncounter(userId int, encounter *model.Encounter) (*model.EncounterCompletion, error)
	FinishEncounter(userId int, encounter *model.Encounter) (*model.EncounterCompletion, error)
}
