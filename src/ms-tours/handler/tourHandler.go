package handler

import (
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

	return router
}
