package domain

import (
  "time"
)

type TimerRepository interface {
	Store(timer Timer) (Timer, error)
  Delete(timer Timer) error
	FindByDayItemID(dayItemID string) ([]Timer, error)
	FindByProjectItemID(dayItemID string) ([]Timer, error)
}

type Timer struct {
	ID string `datastore:"-"`
  ProjectItemID string
  DayItemID string
  IsRunning bool
  Duration int
  Start time.Time
  End time.Time
}