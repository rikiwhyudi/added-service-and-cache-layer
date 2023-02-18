package handlers

import (
	dto "backend-api/dto/result"
	singerdto "backend-api/dto/singer"
	service "backend-api/services"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type singerHandler struct {
	singerService service.SingerService
	validation    *validator.Validate
}

func HandlerSingers(singerService service.SingerService) *singerHandler {
	return &singerHandler{singerService, validator.New()}
}

func (h *singerHandler) FindAllSingers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charshet=utf-8")

	findResponse, err := h.singerService.FindAllSingers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.ArtistResult{Status: http.StatusOK, Action: "find-singers", Data: findResponse}
	json.NewEncoder(w).Encode(response)

}

func (h *singerHandler) GetSingerID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charshet=utf-8")

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	getResponse, err := h.singerService.GetSingerID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.ArtistResult{Status: http.StatusOK, Action: "id-singer", Data: getResponse}
	json.NewEncoder(w).Encode(response)

}

func (h *singerHandler) CreateSinger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charshet=utf-8")

	// Get file path from context
	dataContext := r.Context().Value("dataFile")
	filepath, ok := dataContext.(string)

	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: "Error on get file from context"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Get form value
	var (
		category    = r.FormValue("category")
		name        = r.FormValue("name")
		oldStr      = r.FormValue("old")
		startCareer = r.FormValue("start_career")
	)

	old, err := strconv.Atoi(oldStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Invalid old format"}
		json.NewEncoder(w).Encode(response)
		return
	}

	career, err := time.Parse("02-01-2006", startCareer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Invalid start_career format. Use format: dd-mm-yyyy"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Parse request body to SingerRequest
	request := singerdto.SingerRequest{
		Name:        name,
		Old:         old,
		Thumbnail:   filepath,
		Category:    category,
		StartCareer: career,
	}

	// Validate request input using go-playground/validator
	if err := h.validation.Struct(request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Call service to create singer
	createResponse, err := h.singerService.CreateSinger(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := dto.ArtistResult{Status: http.StatusCreated, Action: "create-singer", Data: createResponse}
	json.NewEncoder(w).Encode(response)
}

func (h *singerHandler) UpdateSinger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charshet=utf-8")

	//Get file path from context
	dataContext := r.Context().Value("dataFile")
	filepath := dataContext.(string)

	// Get form value
	var (
		category    = r.FormValue("category")
		name        = r.FormValue("name")
		oldStr      = r.FormValue("old")
		startCareer = r.FormValue("start_career")
	)

	old, _ := strconv.Atoi(oldStr)
	career, _ := time.Parse("02-01-2006", startCareer)
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Parse request body to UpdateSingerRequest
	request := singerdto.UpdateSingerRequest{
		Name:        name,
		Old:         old,
		Thumbnail:   filepath,
		Category:    category,
		StartCareer: career,
	}

	// Validate request input using go-playground/validator
	if err := h.validation.Struct(request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Call service to update singer
	updateResponse, err := h.singerService.UpdateSinger(request, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.ArtistResult{Status: http.StatusOK, Action: "update-singer", Data: updateResponse}
	json.NewEncoder(w).Encode(response)
}

func (h *singerHandler) DeleteSinger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charshet=utf-8")

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	deleteResponse, err := h.singerService.DeleteSinger(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.ArtistResult{Status: http.StatusOK, Action: "delete-singer", Data: deleteResponse}
	json.NewEncoder(w).Encode(response)
}
