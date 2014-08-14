package domain

import (
  "time"
)

type ProjectRepository interface {
	Store(project Project) (Project, error)
  Delete(project Project) error
	Find(bookID string, active bool) ([]Project, error)
  Get(ID string)(Project, error)
}

type Project struct {
	ID string `datastore:"-"`
  BookID string
	Title string
  Count int
  Open int
  Active bool
  Start time.Time
  End time.Time
}