package usecases

import (
	"ms-lil-service/model"
	"sync"
)

type IRoutine interface {
	Init(wgg *sync.WaitGroup)
	PrintSmth() int
	Die() int
}

type IReceiver interface {
	Init(ch <-chan string) // chan <type>
	ReceiveLoop()
}

type ISender interface {
	Init(ch chan<- string)
	Send()
}

type ICRUDService interface {
	Init()
	GetAll() ([]model.Lil, error)
	Get(id int) (*model.Lil, error)
	Create(l *model.Lil) (*model.Lil, error)
	Delete(id int) error
}

type IRandomService interface {
	ICRUDService
	Update(id int, l *model.Lil) (*model.Lil, error)
}