package main

import (
	"log"
	"net/http"
	"os"
	"tutours/soa/ms-tours/app"
	TourRepository "tutours/soa/ms-tours/dataservice/tourRepository"
	"tutours/soa/ms-tours/handler"
	TourService "tutours/soa/ms-tours/usecase/tourService"

	"github.com/gorilla/mux"
)

func main() {
	app.Init()
	client := app.InitDB()
	storeLogger := log.New(os.Stdout, "[patient-store] ", log.LstdFlags)
	app.InsertInfo(client)

	// Tours setup
	/*
		tourRepo := &tour.TourRepository{Cli: client, Logger: storeLogger}
		tourService := &service.TourService{TourRepository: tourRepo}
		tourHandler := &handler.TourHandler{TourService: tourService}
	*/

	// Keypoints setup
	keypointRepo := &TourRepository.KeypointRepository{Cli: client, Logger: storeLogger}
	keypointService := &TourService.KeypointService{KeypointRepository: keypointRepo}
	keypointHandler := &handler.KeypointHandler{KeypointService: keypointService}

	router := mux.NewRouter()

	//app.SetupTourRoutes(router, tourHandler)
	app.SetupKeypointRoutes(router, keypointHandler)

	log.Fatal(http.ListenAndServe(app.Port, router))
}
