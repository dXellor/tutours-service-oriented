package handler

import (
	"encoding/json"
	"io"
	"log"
	"ms-stakeholders/model"
	"ms-stakeholders/service"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type UserHandler struct {
	UserService *service.UserService
}

func (handler *UserHandler) InitRouter(userService *service.UserService) *chi.Mux {
	handler.UserService = userService

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Post("/auth/login", handler.Login)
	router.Post("/stakeholder/validateToken", handler.ValidateToken)

	return router
}

func (handler *UserHandler) Login(writer http.ResponseWriter, req *http.Request) {
	var credentials model.Credentials
	body, err := io.ReadAll(req.Body)

	if err != nil {
		http.Error(writer, "Failed to read request body", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &credentials)
	if err != nil {
		log.Println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	tokens, err := handler.UserService.Login(credentials)
	if err != nil {
		log.Println("Error while creating a new tour")
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	jsonData, err := json.Marshal(tokens)
	if err != nil {
		log.Println("Error while parsing a new tour")
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(jsonData)
}
func (handler *UserHandler) ValidateToken(writer http.ResponseWriter, req *http.Request) {
	tokenString := req.Header.Get("Authorization")

	if len(tokenString) > 7 && strings.ToUpper(tokenString[0:7]) == "BEARER " {
		tokenString = tokenString[7:]
	}
	_, err := handler.UserService.ValidateToken(tokenString)
	if err != nil {
		log.Println("Not valid token")
		http.Error(writer, err.Error(), http.StatusUnauthorized)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusNoContent)
}
