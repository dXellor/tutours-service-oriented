package app

import (
	"tutours/soa/ms-tours/handler"

	"github.com/gorilla/mux"
)

func SetupTourRoutes(router *mux.Router, tourHandler *handler.TourHandler) {

	//tour routes

	/*
		router.HandleFunc("/tours/{tourId}", tourHandler.Get).Methods("GET")
		router.HandleFunc("/tours/create", tourHandler.Create).Methods("POST")
		router.HandleFunc("/tours/update", tourHandler.Update).Methods("POST")
		router.HandleFunc("/tours/delete/{tourId}", tourHandler.Delete).Methods("DELETE")
		router.HandleFunc("/tours/", tourHandler.GetAll).Methods("GET")
		router.HandleFunc("/tours/author/{authorId}", tourHandler.GetAllByAuthor).Methods("GET")

	*/
}

func SetupKeypointRoutes(router *mux.Router, keypointHandler *handler.KeypointHandler) {

	router.HandleFunc("/keyPoint/all", keypointHandler.GetAll).Methods("GET")
	router.HandleFunc("/keyPoint/", keypointHandler.GetAll).Methods("GET")
	router.HandleFunc("/keyPoint/tour/{tourId}", keypointHandler.GetByTour).Methods("GET")
	router.HandleFunc("/keyPoint/", keypointHandler.Create).Methods("POST")
	router.HandleFunc("/keyPoint/{id}", keypointHandler.Update).Methods("POST")
	router.HandleFunc("/keyPoint/{id}", keypointHandler.Delete).Methods("DELETE")

}
