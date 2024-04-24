package dataservice

import (
	"tutours/soa/ms-tours/model"

	"gorm.io/gorm"
)

type ICRUDRepository interface {
	GetAll() ([]model.Tour, error)
	Get(id int) (*model.Tour, error)
	Create(tour *model.Tour) (*model.Tour, error)
	Delete(id int) error
	Update(tour *model.Tour) (*model.Tour, error)
}

type ITourRepository interface {
	ICRUDRepository
	GetByAuthor(authorId int) ([]model.Tour, error)
	GetPublished() ([]model.Tour, error)
	GetPublishedByAuthor(authorId int) ([]model.Tour, error)
}

type IKeypointRepository interface {
	Init(databaseConnection *gorm.DB)
	GetAll() ([]model.Keypoint, error)
	Create(keypoint *model.Keypoint) (model.Keypoint, error)
	Delete(id int) error
	Update(tour *model.Keypoint) (model.Keypoint, error)
	GetByTour(tourId int) ([]model.Keypoint, error)
}
