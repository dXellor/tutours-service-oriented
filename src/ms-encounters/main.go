package main

import (
	"fmt"
	"net/http"

	"tutours/soa/ms-encounters/handler"
)

func main() {
	router := handler.TestRoutes()
	fmt.Println("Encounters micro-service running")
	http.ListenAndServe(":7007", router)
}
