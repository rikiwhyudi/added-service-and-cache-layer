package routes

import (
	"backend-api/cache"
	"backend-api/handlers"
	"backend-api/pkg/mysql"
	"backend-api/pkg/redis"
	"backend-api/repositories"
	service "backend-api/services"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {

	userRepository := repositories.RepositoryAuth(mysql.DB)
	userCache := cache.NewCache(userRepository, redis.RDB)
	userService := service.NewAuthService(userCache)

	h := handlers.HandlerAuth(userService)

	r.HandleFunc("/register", h.Register).Methods("POST")
	r.HandleFunc("/login", h.Login).Methods("POST")

}
