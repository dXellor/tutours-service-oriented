package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tutours/soa/ms-tours/model"
	"tutours/soa/ms-tours/usecase"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type KeypointHandler struct {
	keypointService usecase.IKeypointService
}

func (handler *KeypointHandler) InitRouter(keypointService usecase.IKeypointService) *chi.Mux {
	handler.keypointService = keypointService

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/all", handler.GetAll)
	router.Get("/", handler.GetAll)
	router.Post("/", handler.Create)
	router.Put("/{id}", handler.Update)
	router.Delete("/{id}", handler.Delete)
	router.Get("/tour/{tourId}", handler.GetByTour)

	return router
}

func (handler *KeypointHandler) GetAll(writer http.ResponseWriter, reader *http.Request) {
	keypoints, err := handler.keypointService.GetAll()
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(keypoints)
}

func (handler *KeypointHandler) Create(writer http.ResponseWriter, reader *http.Request) {
	var keypoint model.Keypoint
	err := json.NewDecoder(reader.Body).Decode(&keypoint)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	createdKeypoint, err := handler.keypointService.Create(&keypoint)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(createdKeypoint)
}

func (handler *KeypointHandler) Update(writer http.ResponseWriter, reader *http.Request) {
	var id, _ = strconv.Atoi(chi.URLParam(reader, "id"))
	var keypoint model.Keypoint
	err := json.NewDecoder(reader.Body).Decode(&keypoint)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	updatedKeypoint, err := handler.keypointService.Update(id, &keypoint)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(updatedKeypoint)
}

func (handler *KeypointHandler) Delete(writer http.ResponseWriter, reader *http.Request) {
	var id, _ = strconv.Atoi(chi.URLParam(reader, "id"))
	err := handler.keypointService.Delete(id)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *KeypointHandler) GetByTour(writer http.ResponseWriter, reader *http.Request) {
	var tourId, _ = strconv.Atoi(chi.URLParam(reader, "tourId"))
	keypoints, err := handler.keypointService.GetByTour(tourId)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(keypoints)
}
