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

	r.HandleFunc("/singer", middleware.Auth(middleware.UploadFile(h.CreateSinger))).Methods("POST")
}
