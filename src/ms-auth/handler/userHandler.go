package handler

import (
	"encoding/json"
	"ms-auth/model"
	"ms-auth/usecase"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type UserHandler struct {
	UserService usecase.IUserService
}

func (handler *UserHandler) InitRouter(userService usecase.IUserService) *chi.Mux {
	handler.UserService = userService

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Post("/login", handler.Login)

	return router
}

func (handler *UserHandler) Login(writer http.ResponseWriter, reader *http.Request) {
	var credentials model.Credentials
	err := json.NewDecoder(reader.Body).Decode(&credentials)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	token, err := handler.UserService.Login(credentials)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(token)
}
