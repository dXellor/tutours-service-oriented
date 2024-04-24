package model

import (
	"encoding/json"
	"time"
	"tutours/soa/ms-tours/model/enum"
)

type Tour struct {
	Id               int                 `bson:"_id,omitempty" json:"id"`
	AuthorId         int                 `bson:"authorId,omitempty" json:"authorId"`
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

func (k *Tour) UnmarshalJSON(data []byte) error {

	var temp struct {
		Id               int
		AuthorId         int
		Name             string
		Description      string
		Price            float64
		Duration         int
		Distance         float64
		Difficulty       enum.TourDifficulty
		TransportType    enum.TransportType
		Status           enum.TourStatus
		StatusUpdateTime time.Time
		Tags             []string
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	k.Id = temp.Id
	k.AuthorId = temp.AuthorId
	k.Name = temp.Name
	k.Price = temp.Price
	k.Duration = temp.Duration
	k.Description = temp.Description
	k.Distance = temp.Distance
	k.Difficulty = temp.Difficulty
	k.TransportType = temp.TransportType
	k.Status = temp.Status
	k.StatusUpdateTime = temp.StatusUpdateTime
	k.Tags = temp.Tags

	return nil
}

func (k *Tour) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id               int
		AuthorId         int
		Name             string
		Description      string
		Price            float64
		Duration         int
		Distance         float64
		Difficulty       enum.TourDifficulty
		TransportType    enum.TransportType
		Status           enum.TourStatus
		StatusUpdateTime time.Time
		Tags             []string
	}{
		Id:               k.Id,
		AuthorId:         k.AuthorId,
		Name:             k.Name,
		Price:            k.Price,
		Duration:         k.Duration,
		Description:      k.Description,
		Distance:         k.Distance,
		Difficulty:       k.Difficulty,
		TransportType:    k.TransportType,
		Status:           k.Status,
		StatusUpdateTime: k.StatusUpdateTime,
		Tags:             k.Tags,
	})
}
