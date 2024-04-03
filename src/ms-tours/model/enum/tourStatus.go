package enum

import (
	"encoding/json"
)

type TourStatus int

const (
	DRAFT TourStatus = iota
	PUBLISHED
	ARCHIVED
)

func (ts TourStatus) MarshalJSON() ([]byte, error) {
	switch ts {
	case DRAFT:
		return json.Marshal("DRAFT")
	case PUBLISHED:
		return json.Marshal("PUBLISHED")
	case ARCHIVED:
		return json.Marshal("ARCHIVED")
	default:
		return json.Marshal("DRAFT")
	}
}

func (ts *TourStatus) UnmarshalJSON(b []byte) error {
	s := string(b)
	switch s {
	case "DRAFT":
		*ts = DRAFT
	case "PUBLISHED":
		*ts = PUBLISHED
	case "ARCHIVED":
		*ts = ARCHIVED
	default:
		*ts = DRAFT
	}

	return nil
}
