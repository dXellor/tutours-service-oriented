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
