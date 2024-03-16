package repos

import (
	"fmt"
	"ms-lil-service/model"

	"gorm.io/gorm"
)

type RandomRepo struct {
	DatabaseConnection *gorm.DB
}

func (rr *RandomRepo) Init() {
	fmt.Println("init")
}

func (rr *RandomRepo) GetAll() ([]model.Lil, error) {
	var ls = []model.Lil{}
	dbResult := rr.DatabaseConnection.Find(&ls)
	if dbResult != nil {
		return ls, dbResult.Error
	}
	return ls, nil
}

func (rr *RandomRepo) Get(id int) (model.Lil, error) {
	var l = model.Lil{}
	dbResult := rr.DatabaseConnection.Find(&l, "id = ?", id)
	if dbResult != nil {
		return l, dbResult.Error
	}
	return l, nil
}

func (rr *RandomRepo) Create(l *model.Lil) (model.Lil, error) {
	dbResult := rr.DatabaseConnection.Create(l)
	if dbResult != nil {
		return *l, dbResult.Error
	}
	return *l, nil
}

func (rr *RandomRepo) Update(l *model.Lil) (model.Lil, error) {
	dbResult := rr.DatabaseConnection.Save(l)
	if dbResult != nil {
		return *l, dbResult.Error
	}
	return *l, nil
}

func (rr *RandomRepo) Delete(id int) error {
	dbResult := rr.DatabaseConnection.Delete(&model.Lil{}, id)
	fmt.Println(dbResult)
	if dbResult != nil {
		return dbResult.Error
	}
	return nil
}