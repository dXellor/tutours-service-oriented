package tourservice

import (
	"fmt"
	"tutours/soa/ms-tours/dataservice"
	"tutours/soa/ms-tours/model"
)

type TourService struct {
	TourRepository dataservice.ITourRepository
}

func (tourService *TourService) Init(TourRepository dataservice.ITourRepository) {
	tourService.TourRepository = TourRepository
}

func (tourService *TourService) GetAll() ([]model.Tour, error) {
	tours, err := tourService.TourRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	return tours, nil
}

func (tourService *TourService) Get(id int) (*model.Tour, error) {
	tour, err := tourService.TourRepository.Get(id)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	return tour, nil
}

func (tourService *TourService) Create(tour *model.Tour) (*model.Tour, error) {
	createdTour, err := tourService.TourRepository.Create(tour)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	return createdTour, nil
}

func (tourService *TourService) Update(id int, tour *model.Tour) (*model.Tour, error) {
	tour.Id = id
	li, err := tourService.TourRepository.Update(tour)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	return li, nil
}

func (tourService *TourService) Delete(id int) error {
	err := tourService.TourRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error")
	}
	return nil
}

func (tourService *TourService) GetByAuthor(authorId int) ([]model.Tour, error) {
	fmt.Println("Called byAuthor()")
	tours, error := tourService.TourRepository.GetByAuthor(authorId)
	if error != nil {
		fmt.Println("error")
	}
	return tours, nil
}

func (tourService *TourService) GetPublishedByAuthor(authorId int) ([]model.Tour, error) {
	tours, error := tourService.TourRepository.GetPublishedByAuthor(authorId)
	if error != nil {
		fmt.Println("error")
	}
	return tours, nil
}

func (tourService *TourService) GetPublished() ([]model.Tour, error) {
	tours, error := tourService.TourRepository.GetPublished()
	if error != nil {
		fmt.Println("error")
	}
	return tours, nil
}
