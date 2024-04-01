package TourRepository

import (
	"tutours/soa/ms-tours/model"
	"tutours/soa/ms-tours/model/enum"

	"gorm.io/gorm"
)

type TourRepository struct {
	databaseConnection *gorm.DB
}

func (tourRepository *TourRepository) Init(databaseConnection *gorm.DB) {
	tourRepository.databaseConnection = databaseConnection
}

func (tourRepository *TourRepository) GetAll() ([]model.Tour, error) {
	var tour = []model.Tour{}
	dbResult := tourRepository.databaseConnection.Find(&tour)
	if dbResult != nil {
		return tour, dbResult.Error
	}
	return tour, nil
}

func (tourRepository *TourRepository) Get(id int) (model.Tour, error) {
	var tour = model.Tour{}
	dbResult := tourRepository.databaseConnection.Find(&tour, "\"Id\"=?", id)
	if dbResult != nil {
		return tour, dbResult.Error
	}
	return tour, nil
}

func (tourRepository *TourRepository) Create(tour *model.Tour) (model.Tour, error) {
	dbResult := tourRepository.databaseConnection.Create(tour)
	if dbResult != nil {
		return *tour, dbResult.Error
	}
	return *tour, nil
}

func (tourRepository *TourRepository) Update(tour *model.Tour) (model.Tour, error) {
	dbResult := tourRepository.databaseConnection.Save(tour)
	if dbResult != nil {
		return *tour, dbResult.Error
	}
	return *tour, nil
}

func (tourRepository *TourRepository) Delete(id int) error {
	dbResult := tourRepository.databaseConnection.Delete(&model.Tour{}, id)
	if dbResult != nil {
		return dbResult.Error
	}
	return nil
}

func (tourRepository *TourRepository) GetByAuthor(authorId int) ([]model.Tour, error) {
	var tour = []model.Tour{}
	dbResult := tourRepository.databaseConnection.Find(&tour, "\"UserId\"=?", authorId)
	if dbResult != nil {
		return tour, dbResult.Error
	}
	return tour, nil
}

func (tourRepository *TourRepository) GetPublished() ([]model.Tour, error) {
	var tours []model.Tour
	dbResult := tourRepository.databaseConnection.Find(&tours, "\"Status\"=?", enum.PUBLISHED)
	if dbResult == nil {
		return nil, dbResult.Error
	}

	for i := range tours {
		var keypoints []model.Keypoint
		dbResult := tourRepository.databaseConnection.Where("\"TourId\"=?", tours[i].Id).Find(&keypoints)
		if dbResult == nil {
			return nil, dbResult.Error
		}
		tours[i].Keypoints = keypoints
	}

	return tours, nil
}

func (tourRepository *TourRepository) GetPublishedByAuthor(authorId int) ([]model.Tour, error) {
	var tour = []model.Tour{}
	dbResult := tourRepository.databaseConnection.Find(&tour, "\"Status\"=? and \"UserId\"=?", enum.PUBLISHED, authorId)
	if dbResult != nil {
		return tour, dbResult.Error
	}
	return tour, nil
}
