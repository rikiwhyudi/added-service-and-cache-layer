package routes

import (
	"backend-api/handlers"
	"backend-api/pkg/mysql"
	"backend-api/repositories"
	"backend-api/service"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryAuth(mysql.DB)
	userService := service.NewAuthService(userRepository)
	h := handlers.HandlerAuth(userService)

	r.HandleFunc("/register", h.Register).Methods("POST")
	r.HandleFunc("/login", h.Login).Methods("POST")

}
