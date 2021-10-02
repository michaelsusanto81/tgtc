package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/radityaqb/tgtc/backend/database"
	"github.com/radityaqb/tgtc/backend/handlers"
	"github.com/radityaqb/tgtc/backend/server"
	"github.com/radityaqb/tgtc/backend/service"
)

func main() {

	// Init database connection
	database.InitDB()

	// Init serve HTTP
	router := mux.NewRouter()

	// routes http
	router.HandleFunc("/ping", handlers.Ping).Methods(http.MethodGet)

	// construct your own API endpoints
	// endpoint : /add-product
	// endpoint : /get-product?id=
	// endpoint : /update-product
	// endpoint : /delete-product

	router.HandleFunc("/add-product", handlers.AddProduct).Methods(http.MethodPost)
	router.HandleFunc("/get-product", handlers.GetProduct).Methods(http.MethodGet)
	router.HandleFunc("/update-product", handlers.UpdateProduct).Methods(http.MethodPatch)
	router.HandleFunc("/delete-product", handlers.DeleteProduct).Methods(http.MethodDelete)

	router.HandleFunc("/v2/get-product", service.GetProduct).Methods(http.MethodGet)

	serverConfig := server.Config{
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		Port:         8000,
	}
	server.Serve(serverConfig, router)
}
