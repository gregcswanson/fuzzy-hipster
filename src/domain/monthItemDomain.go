package domain

import (
  "time"
  "errors"
)

type MonthItemRepository interface {
  Store(dayItem MonthItem) (MonthItem, error)
  Delete(id string) error
	Find(year int, month int) ([]MonthItem, error)
  Get(id string) (MonthItem, error)
}

type MonthItem struct {
	ID string `datastore:"-"`
  ProjectID string
  ProjectItemID string
  Month int // stored as MM
  Year int // stored as yyyy
	Status string `datastore:",noindex"`
  Text string
  Sort int64 `datastore:",noindex"`
  Duration int `datastore:",noindex"` // future enhancement to track time spent
  Start time.Time `datastore:",noindex"` // future enhancement to track time spent
  End time.Time `datastore:",noindex"` // future enhancement to track time spent
}

type MonthItems []MonthItem
func (a MonthItems) Find(id string) (int, error) {
  for index, monthItem := range a {
    if monthItem.ID == id {
      return index, nil
    }
	}
  return -1, errors.New("Not Found")
}

type MonthItemBySort []MonthItem
func (a MonthItemBySort) Len() int { return len(a) }
func (a MonthItemBySort) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a MonthItemBySort) Less(i, j int) bool { return a[i].Sort < a[j].Sort }