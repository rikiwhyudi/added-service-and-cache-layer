package routes

import (
	transactionCache "backend-api/cache"
	"backend-api/handlers"
	"backend-api/pkg/middleware"
	"backend-api/pkg/mysql"
	"backend-api/pkg/redis"
	"backend-api/repositories"
	service "backend-api/services"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {

	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	transactionCache := transactionCache.NewTransactionCache(redis.RDB, transactionRepository)
	transactionService := service.NewTransactionService(transactionCache)

	h := handlers.HandlerTransactions(transactionService)

	r.HandleFunc("/transactions", middleware.Auth(h.FindAllTransactions)).Methods("GET")
	r.HandleFunc("/transaction/{id}", middleware.Auth(h.GetTransactionID)).Methods("GET")
	r.HandleFunc("/transaction", middleware.Auth(middleware.UploadFile(h.CreateTransaction))).Methods("POST")
	r.HandleFunc("/transaction/{id}", middleware.Auth(h.UpdateTransaction)).Methods("PATCH")
	r.HandleFunc("/transaction/{id}", middleware.Auth(h.DeleteTransaction)).Methods("DELETE")

}
