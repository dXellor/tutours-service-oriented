package dataservice

import (
	"tutours/soa/ms-tours/model"

	"gorm.io/gorm"
)

type ICRUDRepository interface {
	GetAll() ([]model.Tour, error)
	Get(id int) (model.Tour, error)
	Create(tour *model.Tour) (model.Tour, error)
	Delete(id int) error
	Update(tour *model.Tour) (model.Tour, error)
}

type ITourRepository interface {
	ICRUDRepository
	Init(databaseConnection *gorm.DB)
	GetByAuthor(authorId int) ([]model.Tour, error)
	GetPublished() ([]model.Tour, error)
	GetPublishedByAuthor(authorId int) ([]model.Tour, error)
}
