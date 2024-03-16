package service

import (
	"fmt"
	dataService "ms-lil-service/dataservice"
	"ms-lil-service/model"
)

type RandomService struct {
	Rr dataService.IRandomRepository
}

// should I use this?
func (rs *RandomService) Init() {
	fmt.Println("init");
}

func (rs *RandomService) GetAll() ([]model.Lil, error) {
	ls, err := rs.Rr.GetAll()
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	return ls, nil
}

func (rs *RandomService) Get(id int) (*model.Lil, error) {
	l, err := rs.Rr.Get(id)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	return &l, nil
}

func (rs *RandomService) Create(l *model.Lil) (*model.Lil, error) {
	li, err := rs.Rr.Create(l)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	return &li, nil
}

func (rs *RandomService) Update(id int, l *model.Lil) (*model.Lil, error) {
	l.Id = id
	li, err := rs.Rr.Update(l)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	return &li, nil
}

func (rs *RandomService) Delete(id int) error {
	err := rs.Rr.Delete(id)
	if err != nil {
		return fmt.Errorf("error")
	}
	return nil
}