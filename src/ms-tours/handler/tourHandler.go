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

type TourHandler struct {
	tourservice usecase.ITourService
}

func (handler *TourHandler) InitRouter(tourService usecase.ITourService) *chi.Mux {
	handler.tourservice = tourService

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/all", handler.GetAll)
	router.Get("/", handler.GetAll)
	router.Get("/{id}", handler.Get)
	router.Post("/", handler.Create)
	router.Put("/{id}", handler.Update)
	router.Delete("/{id}", handler.Delete)

	router.Get("/author/{authorId}", handler.Get)
	router.Get("/published", handler.GetPublished)
	router.Get("/published/{authorId}", handler.GetPublishedByAuthor)

	return router
}

func (handler *TourHandler) GetAll(writer http.ResponseWriter, reader *http.Request) {
	tours, err := handler.tourservice.GetAll()
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(tours)
}

func (handler *TourHandler) Get(writer http.ResponseWriter, reader *http.Request) {
	var id, _ = strconv.Atoi(chi.URLParam(reader, "id"))
	tour, err := handler.tourservice.Get(id)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	if tour.Id != 0 {
		json.NewEncoder(writer).Encode(tour)
	}
}

func (handler *TourHandler) Create(writer http.ResponseWriter, reader *http.Request) {
	var tour model.Tour
	err := json.NewDecoder(reader.Body).Decode(&tour)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	createdTour, err := handler.tourservice.Create(&tour)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(createdTour)
}

func (handler *TourHandler) Update(writer http.ResponseWriter, reader *http.Request) {
	var id, _ = strconv.Atoi(chi.URLParam(reader, "id"))
	var tour model.Tour
	err := json.NewDecoder(reader.Body).Decode(&tour)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	updatedTour, err := handler.tourservice.Update(id, &tour)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(updatedTour)
}

func (handler *TourHandler) Delete(writer http.ResponseWriter, reader *http.Request) {
	var id, _ = strconv.Atoi(chi.URLParam(reader, "id"))
	err := handler.tourservice.Delete(id)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *TourHandler) GetPublished(writer http.ResponseWriter, reader *http.Request) {
	tours, err := handler.tourservice.GetPublished()
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(tours)
}

func (handler *TourHandler) GetByAuthor(writer http.ResponseWriter, reader *http.Request) {
	var authorId, _ = strconv.Atoi(chi.URLParam(reader, "authorId"))
	tours, err := handler.tourservice.GetByAuthor(authorId)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(tours)
}

func (handler *TourHandler) GetPublishedByAuthor(writer http.ResponseWriter, reader *http.Request) {
	var authorId, _ = strconv.Atoi(chi.URLParam(reader, "authorId"))

	tours, err := handler.tourservice.GetPublishedByAuthor(authorId)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(tours)
}
