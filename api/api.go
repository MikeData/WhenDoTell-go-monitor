package api

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mikedata/whendotell-go-monitor/mongo"
	"github.com/mikedata/whendotell-go-monitor/tasks"
)

// CreateDatasetAPI manages all the routes configured to API
func CreateAPI(host string, bindAddr string, m mongo.Mongo) {

	// ---- Routing ----
	router := mux.NewRouter()
	api := tasks.TaskAPI{
		DataStore: m,
		Router:    router,
	}

	api.Router.HandleFunc("/", api.Test).Methods("GET")
	api.Router.HandleFunc("/tasks", api.Add).Methods("POST")

	// ---- Server ----
	srv := &http.Server{
		Handler:      api.Router,
		Addr:         bindAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Creating API on: %s", bindAddr)
	log.Fatal(srv.ListenAndServe())

}
