package model

import (
	"time"
	"tutours/soa/ms-encounters/model/enum"
)

type EncounterCompletion struct {
	Id            int `gorm:"primary_key;auto_increment"`
	UserId        int
	LastUpdatedAt time.Time
	EncounterId   int
	Encounter     Encounter
	Xp            int
	Status        enum.EncounterCompletionStatus `json:",string"`
}

func (e *EncounterCompletion) Complete() {
	e.Status = enum.COMPLETED
	e.LastUpdatedAt = time.Now()
}

func (e *EncounterCompletion) Reset() {
	e.Status = enum.STARTED
	e.LastUpdatedAt = time.Now()
}
