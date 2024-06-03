package service

import (
	"errors"
	"fmt"
	"ms-stakeholders/model"
	"ms-stakeholders/repo"
	"ms-stakeholders/util"

	"github.com/dgrijalva/jwt-go"
)

type UserService struct {
	UserRepository *repo.UserRepository
}

func (userService *UserService) Init(userRepository *repo.UserRepository) {
	userService.UserRepository = userRepository
}

var ErrNotFound = errors.New("not found")
var ErrForbidden = errors.New("forbidden")

func (service *UserService) Login(credentials model.Credentials) (model.AuthenticationTokens, error) {
	var tokens model.AuthenticationTokens
	var user model.User
	user, err := service.UserRepository.GetActiveByUsername(credentials.Username)
	if err != nil {
		return tokens, err
	}
	if credentials.Password != user.Password {
		return tokens, fmt.Errorf("Authentication failed: %w", ErrNotFound)
	}
	if user.IsBlocked {
		return tokens, fmt.Errorf("Authentication failed: %w", ErrForbidden)
	}
	if !user.IsEnabled {
		return tokens, fmt.Errorf("Authentication failed: %w", ErrForbidden)
	}

	var person model.Person
	person, err = service.UserRepository.GetPerson(user.Id)
	if err != nil {
		return tokens, err
	}
	fmt.Println(person.Id)
	jwtGen := util.NewJwtGenerator()
	tokens, err = jwtGen.GenerateAccessToken(&user, person.Id)
	if err != nil {
		fmt.Println("Failed to generate access token:", err)
		return tokens, err
	}
	return tokens, nil
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
