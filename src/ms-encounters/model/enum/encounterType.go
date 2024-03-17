package enum

import (
	"encoding/json"
)

type EncounterType int

const (
	SOCIAL EncounterType = iota
	LOCATION
	MISC
)

func (e EncounterType) MarshalJSON() ([]byte, error) {
	switch e {
	case SOCIAL:
		return json.Marshal("SOCIAL")
	case LOCATION:
		return json.Marshal("LOCATION")
	case MISC:
		return json.Marshal("MISC")
	default:
		return json.Marshal("MISC")
	}
}

func (e *EncounterType) UnmarshalJSON(b []byte) error {
	s := string(b)
	switch s {
	case "SOCIAL":
		*e = SOCIAL
	case "LOCATION":
		*e = LOCATION
	case "MISC":
		*e = MISC
	default:
		*e = MISC
	}

	return nil
}