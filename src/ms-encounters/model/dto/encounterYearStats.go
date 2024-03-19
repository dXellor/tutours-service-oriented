package dto

type EncounterYearStats struct {
	Year                   int
	CompletedCountByMonths []int
	FailedCountByMonths    []int
}
