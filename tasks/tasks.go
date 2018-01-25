package tasks

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

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

// Make sure its a valid task and has appropriate accompanying information
func validateTask(task *models.AddTask) (reason string, err error) {

	if task.Task == models.TaskPageHasChanged {

		_, err := url.ParseRequestURI(task.URL)
		return "Failed POST to task. Invalid URL. Must be properly formed, i.e 'http://www.google.com/'", err

	}
	return "Unknown task has passed validation.", errors.New("N/A")
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

	// Validate the task being requested
	issue, err := validateTask(task)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(issue))
		log.Printf(issue)
		return
	}

	api.DataStore.AddTask(task)

}

// Test is a simple test landing page
func (api *TaskAPI) Test(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("We work"))

}
