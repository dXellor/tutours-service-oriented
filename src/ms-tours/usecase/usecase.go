package usecase

import (
	"tutours/soa/ms-tours/dataservice"
	"tutours/soa/ms-tours/model"
)

type ICRUDService interface {
	GetAll() ([]model.Tour, error)
	Get(id int) (*model.Tour, error)
	Create(tour *model.Tour) (*model.Tour, error)
	Delete(id int) error
	Update(id int, tour *model.Tour) (*model.Tour, error)
}

type ITourService interface {
	ICRUDService
	Init(crudRepository dataservice.ITourRepository)
	GetByAuthor(authorId int) ([]model.Tour, error)
	GetPublished() ([]model.Tour, error)
	GetPublishedByAuthor(authorId int) ([]model.Tour, error)
}
