package models

import (
	"time"
)

// A list of reusable resource type across application
const (
	TaskPageHasChanged = "page has changed"
)

// Task represents a resource or "thing to monitor"
type Task struct {
	Name        string    `bson:"name,omitempty"              json:"name,omitempty"`
	Task        string    `bson:"task,omitempty"              json:"task,omitempty"`
	Interval    *Interval `bson:"interval,omitempty"          json:"interval,omitempty"`
	URL         string    `bson:"url,omitempty"               json:"url,omitempty"`
	LastChecked time.Time `bson:"last_checked"                json:"last_checked"`
}

// Interval between running tasks
type Interval struct {
	Hours   int `bson:"hours,omitempty"              json:"hours,omitempty"`
	Minutes int `bson:"minutes,omitempty"            json:"minutes,omitempty"`
}

// OptionsTaskPageHasChanged contains the json options that'll be sent to redis regarding this task
type OptionsTaskPageHasChanged struct {
	URL  string `bson:"url,omitempty"               json:"url,omitempty"`
	Name string `bson:"name,omitempty"              json:"name,omitempty"` // for checking tasks go into the correct queue
}
