package model

type Encounter struct {
	Id             int `json: "-" gorm:"primary_key;auto_increment;"`
	UserId         int
	Name           string
	Description    string
	Latitude       float64
	Longitude      float64
	Xp             int
	Status         EncounterStatus `json:",string"`
	Type           EncounterType   `json:",string"`
	Range          float64
	Image          string
	PeopleCount    int
	ApprovalStatus EncounterApprovalStatus `json:",string"`
	ImageLatitude  float64
	ImageLongitude float64
}

type EncounterType int

const (
	SOCIAL   EncounterType = iota
	LOCATION EncounterType = iota
	MISC     EncounterType = iota
)

type EncounterStatus int

const (
	ACTIVE   EncounterType = iota
	DRAFT    EncounterType = iota
	ARCHIVED EncounterType = iota
)

type EncounterApprovalStatus int

const (
	PENDING         EncounterType = iota
	SYSTEM_APPROVED EncounterType = iota
	ADMIN_APPROVED  EncounterType = iota
	DECLINED        EncounterType = iota
)