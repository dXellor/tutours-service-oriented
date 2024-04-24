package encounterservice

import (
	"fmt"
	"time"
	"tutours/soa/ms-encounters/dataservice"
	"tutours/soa/ms-encounters/model"
	"tutours/soa/ms-encounters/model/enum"
)

type EncounterComlpetionService struct {
	encounterCompletionRepository dataservice.IEncounterCompletionRepository
}

func (encounterCompletionService *EncounterComlpetionService) Init(encounterRepository dataservice.IEncounterCompletionRepository) {
	encounterCompletionService.encounterCompletionRepository = encounterRepository
}

func (encounterCompletionService *EncounterComlpetionService) GetByUser(userId int) ([]model.EncounterCompletion, error) {
	encounters, err := encounterCompletionService.encounterCompletionRepository.GetByUser(userId)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	return encounters, nil
}

func (encounterCompletionService *EncounterComlpetionService) StartEncounter(userId int, encounter *model.Encounter) (*model.EncounterCompletion, error) {

	if !encounterCompletionService.encounterCompletionRepository.HasUserStartedEncounter(userId, encounter.Id) {
		//TODO: Add position condition
		return nil, fmt.Errorf("encounter already started")
	}

	newEncounterCompletition := model.EncounterCompletion{
		UserId:        userId,
		LastUpdatedAt: time.Now(),
		EncounterId:   encounter.Id,
		Xp:            encounter.Xp,
		Status:        enum.STARTED,
	}

	newEncounterCompletition, errUpdate := encounterCompletionService.encounterCompletionRepository.Create(&newEncounterCompletition)
	if errUpdate != nil {
		return nil, fmt.Errorf("Database error:" + errUpdate.Error())
	}
	return &newEncounterCompletition, nil
}

func (encounterCompletionService *EncounterComlpetionService) FinishEncounter(userId int, encounter *model.Encounter) (*model.EncounterCompletion, error) {
	encouterCompletion, err := encounterCompletionService.encounterCompletionRepository.GetByUserAndEncounter(userId, encounter.Id)
	if err != nil {
		return nil, fmt.Errorf("Database error:" + err.Error())
	}

	if encouterCompletion == nil {
		return nil, fmt.Errorf("encounter not started")
	}

	encouterCompletion.Complete()
	encouterCompletionUpdated, errUpdate := encounterCompletionService.encounterCompletionRepository.Update(encouterCompletion)
	if errUpdate != nil {
		return nil, fmt.Errorf("Database error:" + errUpdate.Error())
	}
	return &encouterCompletionUpdated, nil
}
