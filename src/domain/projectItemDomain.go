package domain

import (
  "time"
  "errors"
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

type ProjectItems []ProjectItem
func (a ProjectItems) Find(id string) (int, error) {
  for index, projectItem := range a {
    if projectItem.ID == id {
      return index, nil
    }
	}
  return -1, errors.New("Not Found")
}

type ProjectItemBySort []ProjectItem
func (a ProjectItemBySort) Len() int { return len(a) }
func (a ProjectItemBySort) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ProjectItemBySort) Less(i, j int) bool { return a[i].Sort < a[j].Sort }