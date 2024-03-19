package enum

import (
	"encoding/json"
)

type EncounterCompletionStatus int

const (
	STARTED EncounterCompletionStatus = iota
	FAILED
	COMPLETED
	PROGRESSING
)

func (e EncounterCompletionStatus) MarshalJSON() ([]byte, error) {
	switch e {
	case STARTED:
		return json.Marshal("STARTED")
	case FAILED:
		return json.Marshal("FAILED")
	case COMPLETED:
		return json.Marshal("COMPLETED")
	default:
		return json.Marshal("PROGRESSING")
	}
}

func (e *EncounterCompletionStatus) UnmarshalJSON(b []byte) error {
	s := string(b)
	switch s {
	case "STARTED":
		*e = STARTED
	case "FAILED":
		*e = FAILED
	case "COMPLETED":
		*e = COMPLETED
	default:
		*e = PROGRESSING
	}

	return nil
}
