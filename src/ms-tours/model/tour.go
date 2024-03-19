package model

import (
	"time"
	"tutours/soa/ms-tours/model/enum"
)

type Tour struct {
	Id               int `gorm:"primary_key;auto_increment"`
	UserId           int
	Name             string `gorm:"default:null"`
	Description      string `gorm:"default:null"`
	Price            float64
	Duration         int
	Distance         float64
	TourDifficulty   enum.TourDifficulty `json:",string"`
	TransportType    enum.TransportType  `json:",string"`
	TourStatus       enum.TourStatus     `json:",string"`
	StatusUpdateTime time.Time
	Tags             []string
}
