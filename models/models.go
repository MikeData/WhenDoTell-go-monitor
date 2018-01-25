package models

import (
  "time"
  )

// Task represents a resource or "thing to monitor"
type Task struct {
  Name           string               `bson:"name,omitempty"              json:"name,omitempty"`
  Task           string               `bson:"task,omitempty"              json:"task,omitempty"`
  Interval       *Interval            `bson:"interval,omitempty"          json:"interval,omitempty"`
  URL            string               `bson:"url,omitempty"               json:"url,omitempty"`
  LastChecked    time.Time           `bson:"last_checked"                json:"last_checked"`
}

// Interval between running tasks
type Interval struct {
  Hours          int                  `bson:"hours,omitempty"              json:"hours,omitempty"`
  Minutes        int                  `bson:"minutes,omitempty"            json:"minutes,omitempty"`
}
