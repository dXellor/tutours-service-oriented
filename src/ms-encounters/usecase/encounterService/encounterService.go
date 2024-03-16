package encounterservice

import (
	"fmt"
	"tutours/soa/ms-encounters/dataservice"
	"tutours/soa/ms-encounters/model"
)

type EncounterService struct {
	encounterRepository dataservice.IEncounterRepository
}

func (encounterService *EncounterService) Init(encounterRepository dataservice.IEncounterRepository) {
	encounterService.encounterRepository = encounterRepository
}

func (encounterService *EncounterService) GetAll() ([]model.Encounter, error) {
	encounters, err := encounterService.encounterRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	return encounters, nil
}

func (encounterService *EncounterService) Get(id int) (*model.Encounter, error) {
	encounter, err := encounterService.encounterRepository.Get(id)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	return &encounter, nil
}

func (encounterService *EncounterService) Create(encounter *model.Encounter) (*model.Encounter, error) {
	createdEncounter, err := encounterService.encounterRepository.Create(encounter)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	return &createdEncounter, nil
}

func (encounterService *EncounterService) Update(id int, encounter *model.Encounter) (*model.Encounter, error) {
	encounter.Id = id
	li, err := encounterService.encounterRepository.Update(encounter)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	return &li, nil
}

func (encounterService *EncounterService) Delete(id int) error {
	err := encounterService.encounterRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error")
	}
	return nil
}