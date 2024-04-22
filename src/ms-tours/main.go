package main

import (
	"fmt"
	"net/http"
	"tutours/soa/ms-tours/handler"
	tourservice "tutours/soa/ms-tours/usecase/tourService"
)

func main() {
	
	tourService := tourservice.TourService{}

	tourhandler := handler.TourHandler{}

	router := tourhandler.InitRouter(&tourService)

	fmt.Println("Tours micro-service running")
	http.ListenAndServe(":7007", router)

}
