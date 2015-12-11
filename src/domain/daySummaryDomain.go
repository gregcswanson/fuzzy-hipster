package domain

import (
  "errors"
)

type DaySummaryRepository interface {
  	Store(daySummary DaySummary) (DaySummary, error)
  	Delete(id string) error
  	FindForDay(day int) (DaySummary, error)
	FindForMonth(year int, month int) ([]DaySummary, error)
  	Get(id string) (DaySummary, error)
}

type DaySummary struct {
	ID string `datastore:"-"`
  	Day int // stored as yyyyMMdd
	DayOfMonth int // stored as d
  	Month int // stored as M
  	Year int // stored as yyyy
	Text string
  	Duration int `datastore:",noindex"` // future enhancement to track time spent
}

type DaySummaries []DaySummary
func (a DaySummaries) Find(id string) (int, error) {
	for index, daySummary := range a {
    	if daySummary.ID == id {
      		return index, nil
    	}
	}
  	return -1, errors.New("Not Found")
}

type DaySummaryBySort []DaySummary
func (a DaySummaryBySort) Len() int { return len(a) }
func (a DaySummaryBySort) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a DaySummaryBySort) Less(i, j int) bool { return a[i].Day < a[j].Day }