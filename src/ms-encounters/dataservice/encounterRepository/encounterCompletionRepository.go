package encounterrepository

import (
	"tutours/soa/ms-encounters/model"
	"tutours/soa/ms-encounters/model/enum"

	"gorm.io/gorm"
)

type EncounterCompletionRepository struct {
	databaseConnection *gorm.DB
}

func (encounterCompletionRepository *EncounterCompletionRepository) Init(databaseConnection *gorm.DB) {
	encounterCompletionRepository.databaseConnection = databaseConnection
}

func (encounterCompletionRepository *EncounterCompletionRepository) GetByUser(userId int) ([]model.EncounterCompletion, error) {
	var encounterCompletions = []model.EncounterCompletion{}
	dbResult := encounterCompletionRepository.databaseConnection.Preload("Encounter").Find(&encounterCompletions, "\"UserId\"=?", userId)
	if dbResult != nil {
		return encounterCompletions, dbResult.Error
	}
	return encounterCompletions, nil
}

func (encounterCompletionRepository *EncounterCompletionRepository) Create(encounterCompletion *model.EncounterCompletion) (model.EncounterCompletion, error) {
	dbResult := encounterCompletionRepository.databaseConnection.Create(encounterCompletion)
	if dbResult != nil {
		return *encounterCompletion, dbResult.Error
	}
	return *encounterCompletion, nil
}

func (encounterCompletionRepository *EncounterCompletionRepository) Update(encounterCompletion *model.EncounterCompletion) (model.EncounterCompletion, error) {
	dbResult := encounterCompletionRepository.databaseConnection.Save(encounterCompletion)
	if dbResult != nil {
		return *encounterCompletion, dbResult.Error
	}
	return *encounterCompletion, nil
}

func (encounterCompletionRepository *EncounterCompletionRepository) HasUserStartedEncounter(userId int, encounterId int) bool {
	encounterCompletion, _ := encounterCompletionRepository.GetByUserAndEncounter(userId, encounterId)
	return encounterCompletion != nil
}

func (encounterCompletionRepository *EncounterCompletionRepository) GetByUserAndEncounter(userId int, encounterId int) (*model.EncounterCompletion, error) {
	var encounterCompletion = model.EncounterCompletion{}
	dbResult := encounterCompletionRepository.databaseConnection.Find(&encounterCompletion, "\"UserId\"=? AND \"EncounterId\"=?", userId, encounterId)
	if dbResult != nil {
		return &encounterCompletion, dbResult.Error
	}
	return &encounterCompletion, nil
}

// Statistics
func (encounterCompletionRepository *EncounterCompletionRepository) GetCompletedCountByUser(userId int) (int64, error) {
	var completedCount int64
	var encounterCompletion = []model.EncounterCompletion{}
	dbResult := encounterCompletionRepository.databaseConnection.Find(&encounterCompletion, "\"UserId\"=? AND \"Status\"=?", userId, enum.COMPLETED).Count(&completedCount)
	if dbResult != nil {
		return completedCount, dbResult.Error
	}
	return completedCount, nil
}

func (encounterCompletionRepository *EncounterCompletionRepository) GetFailedCountByUser(userId int) (int64, error) {
	var failedCount int64
	var encounterCompletion = []model.EncounterCompletion{}
	dbResult := encounterCompletionRepository.databaseConnection.Find(&encounterCompletion, "\"UserId\"=? AND \"Status\"=?", userId, enum.FAILED).Count(&failedCount)
	if dbResult != nil {
		return failedCount, dbResult.Error
	}
	return failedCount, nil
}

func (encounterCompletionRepository *EncounterCompletionRepository) GetCompletedCountByUserAndMonth(userId int, month int, year int) (int64, error) {
	var completedCount int64
	var encounterCompletion = []model.EncounterCompletion{}
	dbResult := encounterCompletionRepository.databaseConnection.Find(&encounterCompletion, "\"UserId\"=? AND \"Status\"=? AND DATE_PART('MONTH', \"LastUpdatedAt\")=? AND DATE_PART('YEAR', \"LastUpdatedAt\")=?", userId, enum.COMPLETED, month, year).Count(&completedCount)
	if dbResult != nil {
		return completedCount, dbResult.Error
	}
	return completedCount, nil
}

func (encounterCompletionRepository *EncounterCompletionRepository) GetFailedCountByUserAndMonth(userId int, month int, year int) (int64, error) {
	var failedCount int64
	var encounterCompletion = []model.EncounterCompletion{}
	dbResult := encounterCompletionRepository.databaseConnection.Find(&encounterCompletion, "\"UserId\"=? AND \"Status\"=? AND DATE_PART('MONTH', \"LastUpdatedAt\")=? AND DATE_PART('YEAR', \"LastUpdatedAt\")=?", userId, enum.FAILED, month, year).Count(&failedCount)
	if dbResult != nil {
		return failedCount, dbResult.Error
	}
	return failedCount, nil
}
