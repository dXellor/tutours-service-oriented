package service

import (
	"errors"
	"fmt"
	"ms-auth/dataservice"
	"ms-auth/model"
	"ms-auth/util"

	"github.com/dgrijalva/jwt-go"
)

type UserService struct {
	userRepository dataservice.IUserRepository
}

func (userService *UserService) Init(userRepository dataservice.IUserRepository) {
	userService.userRepository = userRepository
}

var ErrNotFound = errors.New("not found")
var ErrForbidden = errors.New("forbidden")

func (service *UserService) Login(credentials model.Credentials) (model.AuthenticationToken, error) {
	var token model.AuthenticationToken
	var user model.User
	user, err := service.userRepository.GetActiveByUsername(credentials.Username)
	if err != nil {
		return token, err
	}
	if credentials.Password != user.Password {
		return token, fmt.Errorf("authentication failed: %w", ErrNotFound)
	}
	if user.IsBlocked {
		return token, fmt.Errorf("authentication failed: %w", ErrForbidden)
	}
	if !user.IsEnabled {
		return token, fmt.Errorf("authentication failed: %w", ErrForbidden)
	}

	var person model.Person
	person, err = service.userRepository.GetPerson(user.Id)
	if err != nil {
		return token, err
	}

	jwtGen := util.NewJwtGenerator()
	token, err = jwtGen.GenerateAccessToken(&user, person.Id)
	if err != nil {
		fmt.Println("Failed to generate access token:", err)
		return token, err
	}
	return token, nil
}

func (service *UserService) ValidateToken(token string) (jwt.Claims, error) {
	var claims jwt.MapClaims
	jwtGen := util.NewJwtGenerator()
	claims, err := jwtGen.ValidateAccessToken(token)
	if err != nil {
		fmt.Println("Not valid token:", err)
		return claims, err
	}
	return claims, nil
}
