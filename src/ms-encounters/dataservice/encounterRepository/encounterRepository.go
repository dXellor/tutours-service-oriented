package encounterrepository

import (
	"tutours/soa/ms-encounters/model"

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
	dbResult := encounterRepository.databaseConnection.Find(&encounter, "id = ?", id)
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