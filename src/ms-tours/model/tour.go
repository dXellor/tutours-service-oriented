package model

import (
	"database/sql/driver"
	"strings"
	"time"
	"tutours/soa/ms-tours/model/enum"
)

type Tour struct {
	Id               int `gorm:"primary_key;auto_increment"`
	UserId           int
	Name             string
	Description      string
	Price            float64
	Duration         int
	Distance         float64
	Difficulty       enum.TourDifficulty `json:",string"`
	TransportType    enum.TransportType  `json:",string"`
	Status           enum.TourStatus     `json:",string"`
	StatusUpdateTime time.Time
	Tags             []string   `gorm:"-"`
	Keypoints        []Keypoint // New field to hold keypoints
}

func (tour *Tour) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	str, ok := value.(string)
	if !ok {
		return nil
	}
	tags := strings.Split(strings.Trim(str, "{}"), ",")
	tour.Tags = tags
	return nil
}

func (tour Tour) Value() (driver.Value, error) {
	if len(tour.Tags) == 0 {
		return nil, nil
	}
	return "{" + strings.Join(tour.Tags, ",") + "}", nil
}
