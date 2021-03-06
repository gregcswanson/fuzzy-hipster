package domain

import (
  "time"
  "errors"
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
  Sort int64 `datastore:",noindex"`
  Duration int `datastore:",noindex"` // future enhancement to track time spent
  Start time.Time `datastore:",noindex"` // future enhancement to track time spent
  End time.Time `datastore:",noindex"` // future enhancement to track time spent
}

type DayItems []DayItem
func (a DayItems) Find(id string) (int, error) {
  for index, dayItem := range a {
    if dayItem.ID == id {
      return index, nil
    }
	}
  return -1, errors.New("Not Found")
}

type DayItemBySort []DayItem
func (a DayItemBySort) Len() int { return len(a) }
func (a DayItemBySort) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a DayItemBySort) Less(i, j int) bool { return a[i].Sort < a[j].Sort }