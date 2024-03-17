package encounterrepository

import (
	"tutours/soa/ms-encounters/model"
	"tutours/soa/ms-encounters/model/enum"

	"gorm.io/gorm"
)

type EncounterRepository struct {
	databaseConnection *gorm.DB
}

func (encounterRepository *EncounterRepository) Init(databaseConnection *gorm.DB) {
	encounterRepository.databaseConnection = databaseConnection
}

func (encounterRepository *EncounterRepository) GetAll() ([]model.Encounter, error) {
	var encounter = []model.Encounter{}
	dbResult := encounterRepository.databaseConnection.Find(&encounter)
	if dbResult != nil {
		return encounter, dbResult.Error
	}
	return encounter, nil
}

func (encounterRepository *EncounterRepository) Get(id int) (model.Encounter, error) {
	var encounter = model.Encounter{}
	dbResult := encounterRepository.databaseConnection.Find(&encounter, "\"Id\"=?", id)
	if dbResult != nil {
		return encounter, dbResult.Error
	}
	return encounter, nil
}

func (encounterRepository *EncounterRepository) Create(encounter *model.Encounter) (model.Encounter, error) {
	dbResult := encounterRepository.databaseConnection.Create(encounter)
	if dbResult != nil {
		return *encounter, dbResult.Error
	}
	return *encounter, nil
}

func (encounterRepository *EncounterRepository) Update(encounter *model.Encounter) (model.Encounter, error) {
	dbResult := encounterRepository.databaseConnection.Save(encounter)
	if dbResult != nil {
		return *encounter, dbResult.Error
	}
	return *encounter, nil
}

func (encounterRepository *EncounterRepository) Delete(id int) error {
	dbResult := encounterRepository.databaseConnection.Delete(&model.Encounter{}, id)
	if dbResult != nil {
		return dbResult.Error
	}
	return nil
}

func (encounterRepository *EncounterRepository) GetApprovedByStatus(status enum.EncounterStatus) ([]model.Encounter, error) {
	var encounter = []model.Encounter{}
	dbResult := encounterRepository.databaseConnection.Find(&encounter, "\"Status\"=?", status)
	if dbResult != nil {
		return encounter, dbResult.Error
	}
	return encounter, nil
}

func (encounterRepository *EncounterRepository) GetByUser(userId int) ([]model.Encounter, error) {
	var encounter = []model.Encounter{}
	dbResult := encounterRepository.databaseConnection.Find(&encounter, "\"UserId\"=?", userId)
	if dbResult != nil {
		return encounter, dbResult.Error
	}
	return encounter, nil
}

func (encounterRepository *EncounterRepository) GetTouristCreatedEncounters() ([]model.Encounter, error) {
	var encounter = []model.Encounter{}
	dbResult := encounterRepository.databaseConnection.Find(&encounter, "\"ApprovalStatus\" !=?", enum.SYSTEM_APPROVED)
	if dbResult != nil {
		return encounter, dbResult.Error
	}
	return encounter, nil
}