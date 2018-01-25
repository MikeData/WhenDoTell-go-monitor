package api

import (
  "github.com/gorilla/mux"
  "github.com/mikedata/go-data-source-monitor/mongo"
  "github.com/mikedata/go-data-source-monitor/tasks"
  "log"
  "net/http"
  "time"
)



// CreateDatasetAPI manages all the routes configured to API
func CreateAPI(host string, bindAddr string, m mongo.Mongo) {


  // ---- Routing ----
  router := mux.NewRouter()
  api := tasks.TaskAPI{
		DataStore:          m,
		Router:             router,
	}

  api.Router.HandleFunc("/", api.Test).Methods("GET")
  api.Router.HandleFunc("/tasks", api.Add).Methods("POST")


  // ---- Server ----
  srv := &http.Server{
        Handler:      api.Router,
        Addr:         bindAddr,
        // Good practice: enforce timeouts for servers you create!
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

  log.Printf("Creating API on: %s", bindAddr)
  log.Fatal(srv.ListenAndServe())

}
