package handlers

import (
	musicdto "backend-api/dto/music"
	dto "backend-api/dto/result"
	service "backend-api/services"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type musicHandler struct {
	musicService service.MusicService
	validation   *validator.Validate
}

func HandlerMusics(musicService service.MusicService) *musicHandler {
	return &musicHandler{musicService, validator.New()}
}

func (h *musicHandler) FindAllMusics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charshet=utf-8")

	findResponse, err := h.musicService.FindAllMusics()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.MusicResult{Status: http.StatusOK, Action: "find-musics", Data: findResponse}
	json.NewEncoder(w).Encode(response)
}

func (h *musicHandler) GetMusicID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charshet=utf-8")

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	getResponse, err := h.musicService.GetMusicID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.MusicResult{Status: http.StatusOK, Action: "id-music", Data: getResponse}
	json.NewEncoder(w).Encode(response)
}

func (h *musicHandler) CreateMusic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charshet=utf-8")

	// dataContext := r.Context().Value("dataFile")
	// filepath, ok := dataContext.(string)

	// if !ok {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: "Error on get file from context"}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }

	// Get the form value
	var (
		title    = r.FormValue("title")
		years    = r.FormValue("year")
		singerID = r.FormValue("singer_id")
	)

	idSinger, err := strconv.Atoi(singerID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	year, err := time.Parse("02-01-2006", years)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Invalid release format. Use format: dd-mm-yyyy"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Parse request body to MusicRequest
	request := musicdto.MusicRequest{
		Title: title,
		// Thumbnail: filepath,
		Year:     year,
		SingerID: idSinger,
		// MusicFile: filepath,
	}

	// Validate request input using go-playground/validator
	if err := h.validation.Struct(request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Call service to create music
	createResponse, err := h.musicService.CreateMusic(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := dto.MusicResult{Status: http.StatusCreated, Action: "create-music", Data: createResponse}
	json.NewEncoder(w).Encode(response)

}

func (h *musicHandler) UpdateMusic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// dataContext := r.Context().Value("dataFile")
	// filepath := dataContext.(string)

	// musicContext := r.Context().Value("musicFile")
	// filepathmp3 := musicContext.(string)

	// Get the form value
	var (
		title = r.FormValue("title")
		years = r.FormValue("years")
		// singerID = r.FormValue("singer_id")
	)

	// idSinger, err := strconv.Atoi(singerID)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }

	year, err := time.Parse("02-01-2006", years)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Invalid release format. Use format: dd-mm-yyyy"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Parse request body to UpdateMusicRequest
	request := musicdto.UpdatedMusicRequest{
		Title: title,
		// Thumbnail: filepath,
		Year: year,
		// SingerID: idSinger,
		// MusicFile: filepathmp3,
	}

	// Validate request input using go-playground/validator
	if err := h.validation.Struct(request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Call service to update music
	updateResponse, err := h.musicService.UpdateMusic(request, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.MusicResult{Status: http.StatusOK, Action: "update-music", Data: updateResponse}
	json.NewEncoder(w).Encode(response)
}

func (h *musicHandler) DeleteMusic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	deleteResponse, err := h.musicService.DeleteMusic(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.MusicResult{Status: http.StatusOK, Action: "delete-music", Data: deleteResponse}
	json.NewEncoder(w).Encode(response)

}
