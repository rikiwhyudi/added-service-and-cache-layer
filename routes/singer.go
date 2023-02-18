package routes

import (
	singerCache "backend-api/cache"
	"backend-api/handlers"
	"backend-api/pkg/middleware"
	"backend-api/pkg/mysql"
	"backend-api/pkg/redis"
	"backend-api/repositories"
	service "backend-api/services"

	"github.com/gorilla/mux"
)

func SingerRoutes(r *mux.Router) {

	singerRepository := repositories.RepositorySinger(mysql.DB)
	singerCache := singerCache.NewSingerCache(singerRepository, redis.RDB)
	singerService := service.NewSingerService(singerCache)

	h := handlers.HandlerSingers(singerService)

	r.HandleFunc("/singers", middleware.Auth(h.FindAllSingers)).Methods("GET")
	r.HandleFunc("/singer/{id}", middleware.Auth(h.GetSingerID)).Methods("GET")
	r.HandleFunc("/singer", middleware.Auth(middleware.UploadFile(h.CreateSinger))).Methods("POST")
	r.HandleFunc("/singer/{id}", middleware.Auth(middleware.UploadFile(h.UpdateSinger))).Methods("PATCH")
	r.HandleFunc("/singer/{id}", middleware.Auth(h.DeleteSinger)).Methods("DELETE")
}
