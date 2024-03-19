package tourservice

import (
	"tutours/soa/ms-tours/dataservice"
)

type TourService struct {
	tourRepository dataservice.ITourRepository
}

func (tourService *TourService) Init(tourRepository dataservice.ITourRepository) {
	tourService.tourRepository = tourRepository
}
