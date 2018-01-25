package tasks

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mikedata/go-data-source-monitor/models"
	"github.com/mikedata/go-data-source-monitor/mongo"
)

// TaskAPI for attaching api methods to, populated initally in api.go
type TaskAPI struct {
	DataStore     mongo.Mongo
	InternalToken string
	Router        *mux.Router
}

// Add lets you PUT a single task via the API
func (api *TaskAPI) Add(w http.ResponseWriter, r *http.Request) {

	// Read and unmarshal request body
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error reading response.body.", err, nil)
	}

	var task *models.AddTask

	err = json.Unmarshal(bytes, &task)
	if err != nil {
		log.Fatal("Failing to model models.AddTask resource based on request", err)
	}

	api.DataStore.AddTask(task)

}

// Test is a simple test landing page
func (api *TaskAPI) Test(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("We work"))

}
