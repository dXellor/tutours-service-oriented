package model

type Encounter struct {
	Id             int `gorm:"primary_key;auto_increment;"`
	UserId         int
	Name           string
	Description    string
	Latitude       float64
	Longitude      float64
	Xp             int
	Status         int
	Type           EncounterType
	Range          float64
	Image          string
	PeopleCount    int
	ApprovalStatus int
	ImageLatitude  float64
	ImageLongitude float64
}

type EncounterType int

const (
	SOCIAL   EncounterType = iota
	LOCATION EncounterType = iota
	MISC     EncounterType = iota
)