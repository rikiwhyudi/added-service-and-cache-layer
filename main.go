package main

import (
	"backend-api/database"
	"backend-api/pkg/mysql"
	"backend-api/routes"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	// load env
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("failed to load env file!")
	}

	//initial DB connection
	mysql.DatabaseInit()

	r := mux.NewRouter()

	//run migration
	database.RunMigration()

	//grouping routes
	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	//setup static prefix path

	//create cors

	port := os.Getenv("PORT")
	fmt.Println("server running on port " + port)
	//run server
	http.ListenAndServe(":"+port, r)

}