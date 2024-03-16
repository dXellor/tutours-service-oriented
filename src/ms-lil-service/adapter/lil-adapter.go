package adapter

import (
	"encoding/json"
	"fmt"
	"ms-lil-service/model"
	usecases "ms-lil-service/usecase"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type LilAdapter struct {
	Rs usecases.IRandomService
}

func (handler *LilAdapter) InitRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/all", handler.GetAll)
	router.Get("/{id}", handler.Get)
	router.Post("/", handler.Create)
	router.Put("/{id}", handler.Update)
	router.Delete("/{id}", handler.Delete)

	return router
}

func (handler *LilAdapter) GetAll(writer http.ResponseWriter, reader *http.Request) {
	lis, err := handler.Rs.GetAll()
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(lis)
}

func (handler *LilAdapter) GetSimple(writer http.ResponseWriter, reader *http.Request) {
	fmt.Println("get")
	l := model.Lil {
		Name: "test",
		Nickname: "nick",
		Age: 10,
	}
	json.NewEncoder(writer).Encode(l)
}

func (handler *LilAdapter) Get(writer http.ResponseWriter, reader *http.Request) {
	var id, _ = strconv.Atoi(chi.URLParam(reader, "id"))
	li, err := handler.Rs.Get(id)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	if(li.Id != 0) { // find better way, lazy
		json.NewEncoder(writer).Encode(li)
	}
}

func (handler *LilAdapter) Create(writer http.ResponseWriter, reader *http.Request) {
	var l model.Lil
	err := json.NewDecoder(reader.Body).Decode(&l)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	li, err := handler.Rs.Create(&l)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(li)
}

func (handler *LilAdapter) Update(writer http.ResponseWriter, reader *http.Request) {
	var id, _ = strconv.Atoi(chi.URLParam(reader, "id"))
	var l model.Lil
	err := json.NewDecoder(reader.Body).Decode(&l)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	li, err := handler.Rs.Update(id, &l)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(li)
}

func (handler *LilAdapter) Delete(writer http.ResponseWriter, reader *http.Request) {
	var id, _ = strconv.Atoi(chi.URLParam(reader, "id"))
	err := handler.Rs.Delete(id)
	fmt.Println(err)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
}