package routes

import (
	authcache "backend-api/cache"
	"backend-api/handlers"
	"backend-api/pkg/middleware"
	"backend-api/pkg/mysql"
	"backend-api/pkg/redis"
	"backend-api/repositories"
	service "backend-api/services"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {

	userRepository := repositories.RepositoryAuth(mysql.DB)
	userCache := authcache.NewAuthCache(userRepository, redis.RDB)
	userService := service.NewAuthService(userCache)

	h := handlers.HandlerAuth(userService)

	r.HandleFunc("/register", h.Register).Methods("POST")
	r.HandleFunc("/login", h.Login).Methods("POST")
	r.HandleFunc("/check-auth", middleware.Auth(h.GetUserID)).Methods("GET")

}
