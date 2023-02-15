package handlers

import (
	dto "backend-api/dto/result"
	singerdto "backend-api/dto/singer"
	service "backend-api/services"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type singerHandler struct {
	singerService service.SingerService
}

func HandlerSingers(singerService service.SingerService) *singerHandler {
	return &singerHandler{singerService}
}

func (h *singerHandler) CreateSinger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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

	career, err := time.Parse("2006-01-02", startCareer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Invalid start career format"}
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

	// Create singer using SingerService
	createResponse, err := h.singerService.CreateSinger(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := dto.SuccessResult{Status: http.StatusCreated, Action: "create_singer", Data: createResponse}
	json.NewEncoder(w).Encode(response)
}
