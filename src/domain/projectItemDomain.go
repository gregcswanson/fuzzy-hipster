package domain

import (
  "time"
)

type ProjectItemRepository interface {
	Store(projectItem ProjectItem) (ProjectItem, error)
  Delete(projectItem ProjectItem) error
	Find(projectID string) ([]ProjectItem, error)
	Get(id string) (ProjectItem, error)
}

type ProjectItem struct {
	ID string `datastore:"-"`
  ProjectID string
	Status string 
  Text string
  Sort int64
  Duration int `datastore:",noindex"` // future enhancement to track time spent
  Start time.Time `datastore:",noindex"` // future enhancement to track time spent
  End time.Time `datastore:",noindex"` // future enhancement to track time spent
}