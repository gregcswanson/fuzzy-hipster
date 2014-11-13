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
  Description string `datastore:",noindex"`
  Count int `datastore:",noindex"`
  Open int  `datastore:",noindex"`
  Active bool
  Start time.Time `datastore:",noindex"`
  End time.Time `datastore:",noindex"`
}