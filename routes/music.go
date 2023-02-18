package routes

import (
	musicCache "backend-api/cache"
	"backend-api/handlers"
	"backend-api/pkg/middleware"
	"backend-api/pkg/mysql"
	"backend-api/pkg/redis"
	"backend-api/repositories"
	service "backend-api/services"

	"github.com/gorilla/mux"
)

func MusicRoutes(r *mux.Router) {

	musicRepository := repositories.RepositoryMusic(mysql.DB)
	musicCache := musicCache.NewMusicCache(musicRepository, redis.RDB)
	musicService := service.NewMusicService(musicCache)

	h := handlers.HandlerMusics(musicService)

	r.HandleFunc("/musics", middleware.Auth(h.FindAllMusics)).Methods("GET")
	r.HandleFunc("/music/{id}", middleware.Auth(h.GetMusicID)).Methods("GET")
	r.HandleFunc("/music", middleware.Auth(h.CreateMusic)).Methods("POST")
	r.HandleFunc("/music/{id}", middleware.Auth(h.UpdateMusic)).Methods("PATCH")
	r.HandleFunc("/music/{id}", middleware.Auth(h.DeleteMusic)).Methods("DELETE")
}
