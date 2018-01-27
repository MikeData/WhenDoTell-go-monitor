package cron

import (
	"log"

	"time"

	"github.com/go-redis/redis"
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
				log.Print(tasks[i].Name)
				tasks[i].LastChecked = time.Now()
			}

			// sleep for the minimum time measure we're using
			time.Sleep(1 * time.Minute)

		}

	}

}

func makeDuration() {

}
