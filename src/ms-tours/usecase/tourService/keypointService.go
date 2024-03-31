package tourservice

import (
	"fmt"
	"tutours/soa/ms-tours/dataservice"

	"tutours/soa/ms-tours/model"
)

type KeypointService struct {
	keypointRepository dataservice.IKeypointRepository
}

func (keypointService *KeypointService) Init(keypointRepository dataservice.IKeypointRepository) {
	keypointService.keypointRepository = keypointRepository
}

func (keypointService *KeypointService) GetAll() ([]model.Keypoint, error) {
	keypoints, err := keypointService.keypointRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	return keypoints, nil
}

func (keypointService *KeypointService) Create(keypoint *model.Keypoint) (*model.Keypoint, error) {
	createdKeypoint, err := keypointService.keypointRepository.Create(keypoint)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	return &createdKeypoint, nil
}

func (keypointService *KeypointService) Update(id int, keypoint *model.Keypoint) (*model.Keypoint, error) {
	keypoint.Id = id
	li, err := keypointService.keypointRepository.Update(keypoint)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	return &li, nil
}

func (keypointService *KeypointService) Delete(id int) error {
	err := keypointService.keypointRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error")
	}
	return nil
}

func (keypointService *KeypointService) GetByTour(tourId int) ([]model.Keypoint, error) {
	keypoints, error := keypointService.keypointRepository.GetByTour(tourId)
	if error != nil {
		fmt.Println("error")
	}
	return keypoints, nil
}
