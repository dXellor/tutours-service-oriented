package TourRepository

import (
	"tutours/soa/ms-tours/model"

	"gorm.io/gorm"
)

type KeypointRepository struct {
	databaseConnection *gorm.DB
}

func (keypointRepository *KeypointRepository) Init(databaseConnection *gorm.DB) {
	keypointRepository.databaseConnection = databaseConnection
}

func (keypointRepository *KeypointRepository) GetAll() ([]model.Keypoint, error) {
	var keypoints = []model.Keypoint{}
	dbResult := keypointRepository.databaseConnection.Find(&keypoints)
	if dbResult != nil {
		return keypoints, dbResult.Error
	}
	return keypoints, nil
}

func (keypointRepository *KeypointRepository) Create(keypoint *model.Keypoint) (model.Keypoint, error) {
	dbResult := keypointRepository.databaseConnection.Create(keypoint)
	if dbResult != nil {
		return *keypoint, dbResult.Error
	}
	return *keypoint, nil
}

func (keypointRepository *KeypointRepository) Update(keypoint *model.Keypoint) (model.Keypoint, error) {
	dbResult := keypointRepository.databaseConnection.Save(keypoint)
	if dbResult != nil {
		return *keypoint, dbResult.Error
	}
	return *keypoint, nil
}

func (keypointRepository *KeypointRepository) Delete(id int) error {
	dbResult := keypointRepository.databaseConnection.Delete(&model.Keypoint{}, id)
	if dbResult != nil {
		return dbResult.Error
	}
	return nil
}

func (keypointRepository *KeypointRepository) GetByTour(tourId int) ([]model.Keypoint, error) {

	var keypoints = []model.Keypoint{}
	dbResult := keypointRepository.databaseConnection.Find(&keypoints, "\"TourId\"=?", tourId)
	if dbResult.Error != nil {
		return keypoints, dbResult.Error
	}
	return keypoints, nil
}
