package domain

import (
  "time"
)

type DayItemRepository interface {
  Store(dayItem DayItem) (DayItem, error)
  Delete(id string) error
	//FindByProjectID(projectID string) ([]DayItem, error)
  Find(dayAsInt int) ([]DayItem, error)
  FindMonth(year int, month int) ([]DayItem, error)
  Get(id string) (DayItem, error)
}

type DayItem struct {
	ID string `datastore:"-"`
  ProjectID string
  ProjectItemID string
  Day int // stored as yyyyMMdd
	Status string `datastore:",noindex"`
  Text string
  Sort int64
  Duration int `datastore:",noindex"` // future enhancement to track time spent
  Start time.Time `datastore:",noindex"` // future enhancement to track time spent
  End time.Time `datastore:",noindex"` // future enhancement to track time spent
}