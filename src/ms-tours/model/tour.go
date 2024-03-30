package model

import (
	"database/sql/driver"
	"strings"
	"time"
	"tutours/soa/ms-tours/model/enum"
)

// Tour represents a tour model.
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
	Tags             []string `gorm:"-"`
}

func (t *Tour) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	str, ok := value.(string)
	if !ok {
		return nil
	}
	tags := strings.Split(strings.Trim(str, "{}"), ",")
	t.Tags = tags
	return nil
}

func (t Tour) Value() (driver.Value, error) {
	if len(t.Tags) == 0 {
		return nil, nil
	}
	return "{" + strings.Join(t.Tags, ",") + "}", nil
}
