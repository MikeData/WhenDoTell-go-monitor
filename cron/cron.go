package cron

import (
	"log"

	"time"

	"github.com/mikedata/go-data-source-monitor/mongo"
)

// Start monitoring
func Start(m *mongo.Mongo) {

	for {

		// Minus the minimum unit we're measuring by or infinte updates
		currentTime := time.Now().Add(time.Duration(1 * time.Minute))

		tasks, err := m.GetAllTasks()
		if err != nil {
			log.Fatal("Failed to get tasks from Mongo")
		}

		for i := range tasks {
			nextTask := tasks[i].LastChecked.Add(time.Minute * time.Duration(tasks[i].Interval.Minutes))
			if currentTime.After(nextTask) {
				log.Print(tasks[i].Name)
				tasks[i].LastChecked = time.Now()
			}
		}

	}

}

func makeDuration() {

}
