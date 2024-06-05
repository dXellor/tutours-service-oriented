package usecase

import (
	"ms-auth/dataservice"
	"ms-auth/model"
)

type IUserService interface {
	Init(repository dataservice.IUserRepository)
	Login(credentials model.Credentials) (model.AuthenticationToken, error)
}
