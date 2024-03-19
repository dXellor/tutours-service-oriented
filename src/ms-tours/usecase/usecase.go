package usecase

import (
	"tutours/soa/ms-tours/dataservice"
)

type ICRUDService interface {
}

type ITourService interface {
	ICRUDService
	Init(crudRepository dataservice.ITourRepository)
}
