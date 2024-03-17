package enum

import (
	"encoding/json"
)

type EncounterStatus int

const (
	ACTIVE EncounterStatus = iota
	DRAFT
	ARCHIVED
)

func (e EncounterStatus) MarshalJSON() ([]byte, error) {
	switch e {
	case ACTIVE:
		return json.Marshal("ACTIVE")
	case DRAFT:
		return json.Marshal("DRAFT")
	case ARCHIVED:
		return json.Marshal("ARCHIVED")
	default:
		return json.Marshal("ARCHIVED")
	}
}

func (e *EncounterStatus) UnmarshalJSON(b []byte) error {
	s := string(b)
	switch s {
	case "ACTIVE":
		*e = ACTIVE
	case "DRAFT":
		*e = DRAFT
	case "ARCHIVED":
		*e = ARCHIVED
	default:
		*e = ARCHIVED
	}

	return nil
}