package dataService

import "ms-lil-service/model"

type ICRUDRepository interface {
	Init()
	GetAll() ([]model.Lil, error) 
	Get(id int) (model.Lil, error)
	Create(l *model.Lil) (model.Lil, error)
	Delete(id int) error 
}

// just for testing interface inheriting
type IRandomRepository interface {
	ICRUDRepository
	Update(l *model.Lil) (model.Lil, error)
}