package tasks

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/mux"
	"github.com/mikedata/go-data-source-monitor/models"
	"github.com/mikedata/go-data-source-monitor/mongo"
	"github.com/nu7hatch/gouuid"
)

// TaskAPI for attaching api methods to, populated initally in api.go
type TaskAPI struct {
	DataStore     mongo.Mongo
	InternalToken string
	Router        *mux.Router
}

// Make sure its a valid task and has appropriate accompanying information
func validateTask(task *models.Task) error {

	if task.Interval.Hours == 0 && task.Interval.Minutes == 0 {
		return errors.New("You must specify a valid interval to create a new task")
	}

	if task.Task == models.TaskPageHasChanged {

		_, err := url.ParseRequestURI(task.URL)
		if err != nil {
			return errors.New("Failed POST to task. Invalid URL. Must be properly formed, i.e 'http://www.google.com/'")
		}

	} else {
		return errors.New("warning: unknown task has passed validation")
	}

	// everything fine
	return nil
}

// Add lets you PUT a single task via the API
func (api *TaskAPI) Add(w http.ResponseWriter, r *http.Request) {

	// Read and unmarshal request body
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error reading response.body.", err, nil)
	}

	var task *models.Task

	err = json.Unmarshal(bytes, &task)
	if err != nil {
		log.Fatal("Failing to model models.AddTask resource based on request", err)
		w.WriteHeader(400)
		w.Write([]byte("Cannot unmarshall your json request"))
	}

	// Validate the task being requested
	err = validateTask(task)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		log.Print(err)
		return
	}

	// Set last_check as now
	task.LastChecked = time.Now()

	// Create a uuid
	u, err := uuid.NewV4()
	if err != nil {
		log.Print("Failed to generate uuid for new task, aborting task update.")
		w.WriteHeader(404)
		w.Write([]byte(err.Error()))
		return
	}
	task.ID = u.String()

	err = api.DataStore.AddTask(task)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("Error writing to database.")) // TODO crap
		log.Print("Failed write to Mongo.", err)      // TODO crap
		return
	}

	complete := "Successfully created task ''" + task.Task + "' as task type '" + task.Name + "'"
	log.Print(complete)
	w.Write([]byte(complete))

}

// Test is a simple test landing page
func (api *TaskAPI) Test(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("We work"))

}
