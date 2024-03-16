package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tutours/soa/ms-encounters/model"
	"tutours/soa/ms-encounters/usecase"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type EncounterHandler struct {
	encounterService usecase.IEncounterService
}

func (handler *EncounterHandler) InitRouter(encounterService usecase.IEncounterService) *chi.Mux {
	handler.encounterService = encounterService

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

func (handler *EncounterHandler) GetAll(writer http.ResponseWriter, reader *http.Request) {
	encounter, err := handler.encounterService.GetAll()
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(encounter)
}

func (handler *EncounterHandler) Get(writer http.ResponseWriter, reader *http.Request) {
	var id, _ = strconv.Atoi(chi.URLParam(reader, "id"))
	encounter, err := handler.encounterService.Get(id)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	if(encounter.Id != 0) { // find better way, lazy
		json.NewEncoder(writer).Encode(encounter)
	}
}

func (handler *EncounterHandler) Create(writer http.ResponseWriter, reader *http.Request) {
	var encounter model.Encounter
	err := json.NewDecoder(reader.Body).Decode(&encounter)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	createdEncounter, err := handler.encounterService.Create(&encounter)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(createdEncounter)
}

func (handler *EncounterHandler) Update(writer http.ResponseWriter, reader *http.Request) {
	var id, _ = strconv.Atoi(chi.URLParam(reader, "id"))
	var encounter model.Encounter
	err := json.NewDecoder(reader.Body).Decode(&encounter)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	updatedEncounter, err := handler.encounterService.Update(id, &encounter)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(updatedEncounter)
}

func (handler *EncounterHandler) Delete(writer http.ResponseWriter, reader *http.Request) {
	var id, _ = strconv.Atoi(chi.URLParam(reader, "id"))
	err := handler.encounterService.Delete(id)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
}