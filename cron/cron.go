package cron

import (
	"encoding/json"
	"log"

	"time"

	"github.com/go-redis/redis"
	"github.com/mikedata/go-data-source-monitor/models"
	"github.com/mikedata/go-data-source-monitor/mongo"
)

// Start monitoring
func Start(m *mongo.Mongo, r *redis.Client) {

	for {

		currentTime := time.Now()

		tasks, err := m.GetAllTasks()
		if err != nil {
			log.Fatal("Failed to get tasks from Mongo")
		}

		for i := range tasks {
			nextTask := tasks[i].LastChecked.Add(time.Duration(tasks[i].Interval.Minutes) * time.Minute)

			if currentTime.After(nextTask) {

				if tasks[i].Task == models.TaskPageHasChanged {
					produceTaskPageHasChanged(r, tasks[i])
				}

				tasks[i].LastChecked = time.Now()
			}

			// sleep for the minimum time measure we're using
			time.Sleep(1 * time.Minute)

		}

	}

}

func produceTaskPageHasChanged(r *redis.Client, task *models.Task) {

	context := &models.OptionsTaskPageHasChanged{
		URL:  task.URL,
		Name: task.Name,
	}

	details, err := json.Marshal(*context)
	if err != nil {
		log.Print("Failing to unmarshall task for: "+task.Task, err)
		return
	}

	// Send the update to redis
	err = r.Set(task.Task, details, 0).Err()
	if err != nil {
		panic(err)
	}
}
