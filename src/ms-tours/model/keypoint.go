package model

type Keypoint struct {
	Id          int `gorm:"primary_key;auto_increment"`
	TourId      int
	Name        string `gorm:"default:null"`
	Latitude    float64
	Longitude   float64
	Description string `gorm:"default:null"`
	Position    int
	Image       string `gorm:"default:null"`
	Secret      string `gorm:"default:null"`
}
