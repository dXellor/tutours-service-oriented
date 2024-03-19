package tourrepository

import (
	"gorm.io/gorm"
)

type TourRepository struct {
	databaseConnection *gorm.DB
}

func (tourRepository *TourRepository) Init(databaseConnection *gorm.DB) {
	tourRepository.databaseConnection = databaseConnection
}
