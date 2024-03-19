package model

import (
	"time"
	"tutours/soa/ms-encounters/model/enum"
)

type EncounterCompletion struct {
	UserId        int
	LastUpdatedAt time.Time
	EncounterId   int
	Xp            int
	Status        enum.EncounterCompletionStatus
}

func (e *EncounterCompletion) Complete() {
	e.Status = enum.COMPLETED
	e.LastUpdatedAt = time.Now()
}

func (e *EncounterCompletion) Reset() {
	e.Status = enum.STARTED
	e.LastUpdatedAt = time.Now()
}
