package enum

import (
	"encoding/json"
)

type TransportType int

const (
	WALK = iota
	BIKE
	CAR
	BOAT
)

func (tt TransportType) MarshalJSON() ([]byte, error) {
	switch tt {
	case WALK:
		return json.Marshal("WALK")
	case BIKE:
		return json.Marshal("BIKE")
	case CAR:
		return json.Marshal("CAR")
	case BOAT:
		return json.Marshal("BOAT")
	default:
		return json.Marshal("WALK")
	}
}

func (tt *TransportType) UnmarshalJSON(b []byte) error {
	s := string(b)
	switch s {
	case "WALK":
		*tt = WALK
	case "BIKE":
		*tt = BIKE
	case "CAR":
		*tt = CAR
	case "BOAT":
		*tt = BOAT
	default:
		*tt = WALK
	}
	return nil

}
