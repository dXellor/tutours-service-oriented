package app

import (
	"tutours/soa/ms-tours/handler"

	"github.com/gorilla/mux"
)

func SetupTourRoutes(router *mux.Router, tourHandler *handler.TourHandler) {

	router.HandleFunc("/all", tourHandler.GetAll).Methods("GET")
	router.HandleFunc("/", tourHandler.GetAll).Methods("GET")
	router.HandleFunc("/{id}", tourHandler.Get).Methods("GET")
	router.HandleFunc("/{id}", tourHandler.Update).Methods("PUT")
	router.HandleFunc("/{id}", tourHandler.Delete).Methods("DELETE")
	router.HandleFunc("/", tourHandler.Create).Methods("POST")
	router.HandleFunc("/author/{authorId}", tourHandler.GetByAuthor).Methods("GET")
	router.HandleFunc("/published/{authorId}", tourHandler.GetPublishedByAuthor).Methods("GET")
	router.HandleFunc("/published", tourHandler.GetPublished).Methods("GET")

}

func SetupKeypointRoutes(router *mux.Router, keypointHandler *handler.KeypointHandler) {

	router.HandleFunc("/keyPoint/all", keypointHandler.GetAll).Methods("GET")
	router.HandleFunc("/keyPoint/", keypointHandler.GetAll).Methods("GET")
	router.HandleFunc("/keyPoint/tour/{tourId}", keypointHandler.GetByTour).Methods("GET")
	router.HandleFunc("/keyPoint/", keypointHandler.Create).Methods("POST")
	router.HandleFunc("/keyPoint/{id}", keypointHandler.Update).Methods("POST")
	router.HandleFunc("/keyPoint/{id}", keypointHandler.Delete).Methods("DELETE")

}
