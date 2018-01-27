package mongo

import (
	"errors"
	"time"

	"github.com/mikedata/go-data-source-monitor/models"
	"gopkg.in/mgo.v2"
)

// Mongo represents a simplistic MongoDB configuration.
type Mongo struct {
	Collection     string
	Database       string
	APIURL         string
	Session        *mgo.Session
	URI            string
	lastPingTime   time.Time
	lastPingResult error
}

// Init creates a new mgo.Session with a strong consistency and a write mode of "majority".
func (m *Mongo) Init() (session *mgo.Session, err error) {
	if session != nil {
		return nil, errors.New("session already exists")
	}

	if session, err = mgo.Dial(m.URI); err != nil {
		return nil, err
	}

	session.EnsureSafe(&mgo.Safe{WMode: "majority"})
	session.SetMode(mgo.Strong, true)
	return session, nil
}

// AddTask adds a single monitoring task
func (m *Mongo) AddTask(task *models.Task) error {

	s := m.Session.Copy()
	defer s.Close()

	err := s.DB(m.Database).C(m.Collection).Insert(task)
	if err != nil {
		return err
	}

	return nil

}

// GetAllTasks gets all tasks from Mongo
func (m *Mongo) GetAllTasks() ([]*models.Task, error) {

	s := m.Session.Copy()
	defer s.Close()

	var allData []*models.Task
	err := s.DB(m.Database).C(m.Collection).Find(nil).All(&allData)
	if err != nil {
		return nil, err
	}

	return allData, nil
}
