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

	authRepository := repositories.RepositoryAuth(mysql.DB)
	authCache := authcache.NewAuthCache(authRepository, redis.RDB)
	authService := service.NewAuthService(authCache)

	h := handlers.HandlerAuth(authService)

	r.HandleFunc("/register", h.Register).Methods("POST")
	r.HandleFunc("/login", h.Login).Methods("POST")
	r.HandleFunc("/check-auth", middleware.Auth(h.GetUserID)).Methods("GET")

}
