package models


// AddTask represents a resource or "thing to monitor"
type AddTask struct {
  Name           string               `bson:"name,omitempty"              json:"name,omitempty"`
  Task           string               `bson:"task,omitempty"              json:"task,omitempty"`
  Interval       *Interval            `bson:"interval,omitempty"          json:"interval,omitempty"`
  URL            string               `bson:"url,omitempty"               json:"url,omitempty"`
  lastChecked    int64                `bson:"last_checked"`
}


type Interval struct {
  Hours          int                  `bson:"hours,omitempty"              json:"hours,omitempty"`
  Minutes        int                  `bson:"minutes,omitempty"            json:"minutes,omitempty"`
}
