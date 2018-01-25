package models

import (
  "time"
)


// represents a resource or "thing to monitor"
type AddTask struct {
  Name           string            `bson:"name,omitempty"              json:"name,omitempty"`
  Type           string            `bson:"type,omitempty"              json:"type,omitempty"`
  Start          *time.Time        `bson:"start,omitempty"             json:"start,omitempty"`
  Interval       *time.Time        `bson:"interval,omitempty"          json:"interval,omitempty"`
}
