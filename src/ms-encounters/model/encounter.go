package model

import "tutours/soa/ms-encounters/model/enum"

type Encounter struct {
	Id             int `gorm:"primary_key;auto_increment"`
	UserId         int
	Name           string `gorm:"default:null"`
	Description    string `gorm:"default:null"`
	Latitude       float64
	Longitude      float64
	Xp             int
	Status         enum.EncounterStatus `json:",string"`
	Type           enum.EncounterType `json:",string"`
	Range          float64
	Image          string `gorm:"default:null"`
	PeopleCount    int
	ApprovalStatus enum.EncounterApprovalStatus `json:",string"`
	ImageLatitude  float64
	ImageLongitude float64
}