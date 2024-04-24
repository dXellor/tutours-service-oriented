package encounterservice

import (
	"fmt"
	"tutours/soa/ms-encounters/dataservice"
	"tutours/soa/ms-encounters/model/dto"
)

type EncounterStatsService struct {
	encounterCompletionRepository dataservice.IEncounterCompletionRepository
}

func (encounterStatsService *EncounterStatsService) Init(encounterRepository dataservice.IEncounterCompletionRepository) {
	encounterStatsService.encounterCompletionRepository = encounterRepository
}

func (encounterStatsService *EncounterStatsService) GetEncounterStatsByUser(userId int) (*dto.EncounterStats, error) {
	completed, errCompleted := encounterStatsService.encounterCompletionRepository.GetCompletedCountByUser(userId)
	failed, errFailed := encounterStatsService.encounterCompletionRepository.GetFailedCountByUser(userId)
	if errCompleted != nil || errFailed != nil {
		return nil, fmt.Errorf("error")
	}
	stats := dto.EncounterStats{}
	stats.CompletedCount = int(completed)
	stats.FailedCount = int(failed)
	return &stats, nil
}

func (encounterStatsService *EncounterStatsService) GetEncounterYearStatsByUser(userId int, year int) (*dto.EncounterYearStats, error) {
	encounterYearStats := dto.EncounterYearStats{}
	encounterYearStats.Year = year
	for month := 1; month <= 12; month++ {
		var completed, errCompleted = encounterStatsService.encounterCompletionRepository.GetCompletedCountByUserAndMonth(userId, month, year)
		var failed, errFailed = encounterStatsService.encounterCompletionRepository.GetFailedCountByUserAndMonth(userId, month, year)
		if errCompleted == nil {
			completed = 0
		}
		if errFailed == nil {
			failed = 0
		}
		encounterYearStats.CompletedCountByMonths = append(encounterYearStats.CompletedCountByMonths, int(completed))
		encounterYearStats.FailedCountByMonths = append(encounterYearStats.FailedCountByMonths, int(failed))
	}
	return &encounterYearStats, nil
}
