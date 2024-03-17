package enum

import (
	"encoding/json"
)

type EncounterApprovalStatus int

const (
	PENDING EncounterApprovalStatus = iota
	SYSTEM_APPROVED
	ADMIN_APPROVED
	DECLINED
)

func (e EncounterApprovalStatus) MarshalJSON() ([]byte, error) {
	switch e {
	case PENDING:
		return json.Marshal("PENDING")
	case SYSTEM_APPROVED:
		return json.Marshal("SYSTEM_APPROVED")
	case ADMIN_APPROVED:
		return json.Marshal("ADMIN_APPROVED")
	case DECLINED:
		return json.Marshal("DECLINED")
	default:
		return json.Marshal("DECLINED")
	}
}

func (e *EncounterApprovalStatus) UnmarshalJSON(b []byte) error {
	s := string(b)
	switch s {
	case "PENDING":
		*e = PENDING
	case "SYSTEM_APPROVED":
		*e = SYSTEM_APPROVED
	case "ADMIN_APPROVED":
		*e = ADMIN_APPROVED
	case "DECLINED":
		*e = DECLINED
	default:
		*e = DECLINED
	}

	return nil
}