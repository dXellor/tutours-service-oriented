package enum

import (
	"encoding/json"
)

type TourDifficulty int

const (
	EASY = iota
	MEDIUM
	HARD
	EXTREME
)

func (td TourDifficulty) MarshalJSON() ([]byte, error) {
	switch td {
	case EASY:
		return json.Marshal("EASY")
	case MEDIUM:
		return json.Marshal("MEDIUM")
	case HARD:
		return json.Marshal("HARD")
	case EXTREME:
		return json.Marshal("EXTREME")
	default:
		return json.Marshal("EASY")
	}
}

func (td *TourDifficulty) UnmarshalJSON(b []byte) error {
	s := string(b)
	switch s {
	case "EASY":
		*td = EASY
	case "MEDIUM":
		*td = MEDIUM
	case "HARD":
		*td = HARD
	case "EXTREME":
		*td = EXTREME
	default:
		*td = EASY
	}

	return nil
}
