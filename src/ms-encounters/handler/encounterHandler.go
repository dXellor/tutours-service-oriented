package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"tutours/soa/ms-encounters/model"
	"tutours/soa/ms-encounters/model/enum"
	ms_encounters "tutours/soa/ms-encounters/proto"
	"tutours/soa/ms-encounters/usecase"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type EncounterHandler struct {
	encounterService           usecase.IEncounterService
	encounterStatsService      usecase.IEncounterStatsService
	encounterCompletionService usecase.IEncounterCompletionService

	ms_encounters.UnimplementedEncountersServer
}

func (handler *EncounterHandler) InitRouter(encounterService usecase.IEncounterService, encounterStatsService usecase.IEncounterStatsService, encounterCompletionService usecase.IEncounterCompletionService) *chi.Mux {
	handler.encounterService = encounterService
	handler.encounterStatsService = encounterStatsService
	handler.encounterCompletionService = encounterCompletionService

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	//router.Get("/all", handler.GetAll)
	//router.Get("/", handler.GetAll)
	router.Get("/{id}", handler.Get)
	router.Post("/", handler.Create)
	router.Put("/{id}", handler.Update)
	router.Delete("/{id}", handler.Delete)

	router.Get("/status", handler.GetApprovedByStatus)
	router.Get("/byUser/{userId}", handler.GetByUser)
	router.Get("/touristCreatedEncounters", handler.GetTouristCreatedEncounters)

	router.Put("/approve", handler.Approve)
	router.Put("/decline", handler.Decline)

	//Progression
	router.Get("/encounterCompletions/{userId}", handler.GetCompletionsByUser)
	router.Post("/startEncounter/{userId}", handler.StartEncounter)
	router.Put("/finishEncounter/{userId}", handler.FinishEncounter)

	//Statistics
	router.Get("/completions/{userId}", handler.StatsCompletions)
	router.Get("/yearCompletions/{userId}", handler.StatsYearCompletions)

	return router
}

// func (handler *EncounterHandler) GetAll(writer http.ResponseWriter, reader *http.Request) {
// 	encounters, err := handler.encounterService.GetAll()
// 	if err != nil {
// 		writer.WriteHeader(http.StatusExpectationFailed)
// 		return
// 	}
// 	writer.WriteHeader(http.StatusOK)
// 	writer.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(writer).Encode(encounters)
// }

func encounterTypeToGRPC(encounterType enum.EncounterType) ms_encounters.Encounter_EncounterType {
	switch encounterType {
	case enum.SOCIAL:
		return ms_encounters.Encounter_SOCIAL
	case enum.LOCATION:
		return ms_encounters.Encounter_LOCATION
	default:
		return ms_encounters.Encounter_MISC
	}
}
func encounterStatusToGRPC(status enum.EncounterStatus) ms_encounters.EncounterStatus {
	switch status {
	case enum.ACTIVE:
		return ms_encounters.EncounterStatus_ACTIVE
	case enum.DRAFT:
		return ms_encounters.EncounterStatus_DRAFT
	default:
		return ms_encounters.EncounterStatus_ARCHIVED
	}

}
func encounterApprovalStatusToGRPC(status enum.EncounterApprovalStatus) ms_encounters.Encounter_EncounterApprovalStatus {
	switch status {
	case enum.ADMIN_APPROVED:
		return ms_encounters.Encounter_ADMIN_APPROVED
	case enum.SYSTEM_APPROVED:
		return ms_encounters.Encounter_SYSTEM_APPROVED
	case enum.DECLINED:
		return ms_encounters.Encounter_DECLINED
	default:
		return ms_encounters.Encounter_PENDING
	}
}

func encounterToGRPC(e *model.Encounter) *ms_encounters.Encounter {
	return &ms_encounters.Encounter{
		Id:             int32(e.Id),
		UserId:         int32(e.UserId),
		Name:           e.Name,
		Description:    e.Description,
		Latitude:       float32(e.Latitude),
		Longitude:      float32(e.Longitude),
		Xp:             int32(e.Xp),
		Status:         encounterStatusToGRPC(e.Status),
		Type:           encounterTypeToGRPC(e.Type),
		Range:          float32(e.Range),
		Image:          e.Image,
		ImageLatitude:  float32(e.ImageLatitude),
		ImageLongitude: float32(e.ImageLongitude),
		PeopleCount:    int32(e.PeopleCount),
		ApprovalStatus: encounterApprovalStatusToGRPC(e.ApprovalStatus),
	}
}

func encountersToGRPC(encounters []model.Encounter) *ms_encounters.EncountersResponse {
	result := make([]*ms_encounters.Encounter, len(encounters))
	for i, e := range encounters {
		result[i] = encounterToGRPC(&e)
	}
	return &ms_encounters.EncountersResponse{Encounters: result}
}

func (handler *EncounterHandler) GetAll(ctx context.Context, request *ms_encounters.NoParamsRequest) (*ms_encounters.EncountersResponse, error) {
	encounters, err := handler.encounterService.GetAll()
	if err != nil {
		return nil, err
	}
	return encountersToGRPC(encounters), nil
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
	if encounter.Id != 0 { // find better way, lazy
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

func (handler *EncounterHandler) GetApprovedByStatus(writer http.ResponseWriter, reader *http.Request) {
	var statusString = reader.URL.Query().Get("status")
	status := new(enum.EncounterStatus)
	status.UnmarshalJSON([]byte(statusString))
	encounters, err := handler.encounterService.GetApprovedByStatus(*status)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(encounters)
}

func (handler *EncounterHandler) GetByUser(writer http.ResponseWriter, reader *http.Request) {
	var userId, _ = strconv.Atoi(chi.URLParam(reader, "userId"))
	encounters, err := handler.encounterService.GetByUser(userId)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(encounters)
}

func (handler *EncounterHandler) GetTouristCreatedEncounters(writer http.ResponseWriter, reader *http.Request) {
	encounters, err := handler.encounterService.GetTouristCreatedEncounters()
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(encounters)
}

func (handler *EncounterHandler) Approve(writer http.ResponseWriter, reader *http.Request) {
	var encounter model.Encounter
	err := json.NewDecoder(reader.Body).Decode(&encounter)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	updatedEncounter, err := handler.encounterService.Approve(&encounter)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(updatedEncounter)
}

func (handler *EncounterHandler) Decline(writer http.ResponseWriter, reader *http.Request) {
	var encounter model.Encounter
	err := json.NewDecoder(reader.Body).Decode(&encounter)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	updatedEncounter, err := handler.encounterService.Decline(&encounter)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(updatedEncounter)
}

func (handler *EncounterHandler) StatsCompletions(writer http.ResponseWriter, reader *http.Request) {
	var id, _ = strconv.Atoi(chi.URLParam(reader, "userId"))
	encounterStats, err := handler.encounterStatsService.GetEncounterStatsByUser(id)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	fmt.Println(id)
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(encounterStats)
}

func (handler *EncounterHandler) StatsYearCompletions(writer http.ResponseWriter, reader *http.Request) {
	var id, _ = strconv.Atoi(chi.URLParam(reader, "userId"))
	var year, _ = strconv.Atoi(reader.URL.Query().Get("year"))
	encounterYearStats, err := handler.encounterStatsService.GetEncounterYearStatsByUser(id, year)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	fmt.Println(id)
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(encounterYearStats)
}

func (handler *EncounterHandler) GetCompletionsByUser(writer http.ResponseWriter, reader *http.Request) {
	var id, _ = strconv.Atoi(chi.URLParam(reader, "userId"))
	encounterCompletions, err := handler.encounterCompletionService.GetByUser(id)
	if err != nil {
		fmt.Println("End my life")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(encounterCompletions)
}

func (handler *EncounterHandler) StartEncounter(writer http.ResponseWriter, reader *http.Request) {
	var encounter model.Encounter
	err := json.NewDecoder(reader.Body).Decode(&encounter)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	var id, _ = strconv.Atoi(chi.URLParam(reader, "userId"))
	newEncounterCompletition, err := handler.encounterCompletionService.StartEncounter(id, &encounter)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	fmt.Println(id)
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(newEncounterCompletition)
}

func (handler *EncounterHandler) FinishEncounter(writer http.ResponseWriter, reader *http.Request) {
	var encounter model.Encounter
	err := json.NewDecoder(reader.Body).Decode(&encounter)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	var id, _ = strconv.Atoi(chi.URLParam(reader, "userId"))
	encounterCompletion, err := handler.encounterCompletionService.FinishEncounter(id, &encounter)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	fmt.Println(id)
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(encounterCompletion)
}
