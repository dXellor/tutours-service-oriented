package model

import (
	"encoding/json"
	"io"
	"time"
	"tutours/soa/ms-tours/model/enum"
)

type Tour struct {
	Id               int                 `gorm:"primary_key;auto_increment"`
	AuthorId         int                 `gorm:"not null"`
	Name             string              `bson:"name" json:"name"`
	Description      string              `bson:"description" json:"description"`
	Price            float64             `bson:"price" json:"price"`
	Duration         int                 `bson:"duration" json:"duration"`
	Distance         float64             `bson:"distance" json:"distance"`
	Difficulty       enum.TourDifficulty `bson:"difficulty" json:"difficulty"`
	TransportType    enum.TransportType  `bson:"transportType" json:"transportType"`
	Status           enum.TourStatus     `bson:"status" json:"status"`
	StatusUpdateTime time.Time           `bson:"statusUpdateTime" json:"statusUpdateTime"`
	Tags             []string            `bson:"tags,omitempty" json:"tags"`
	Keypoints        []Keypoint
}

type Tours []*Tour

func (p *Tours) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Tour) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Tour) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}
