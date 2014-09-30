package usecases

import (
	"src/interfaces"
  "src/domain"
  "time"
  "errors"
  "log"
)

type DayItem struct {
	ID string
  Day int
  ProjectID string
  ProjectItemID string
  ProjectName string
  Status string // open, closed, note, cancelled, running
	Text string
  Sort int64
}

type DayItemInteractor struct {
	Context interfaces.DomainContext
}

func (interactor *DayItemInteractor) FindByDay(dayAsInt int) ([]DayItem, error) {
	// get the items on the day
  log.Println("usecases.DayItemInteractor.Find")
  foundDayItems, findError := interactor.Context.DayItems.Find(dayAsInt)
  if findError != nil {
    log.Println(findError)
    return []DayItem{}, findError
  }
	// Copy to the use case model
  var dayItems []DayItem
	dayItems = make([]DayItem, len(foundDayItems))
	for i, dayItem := range foundDayItems {
    dayItems[i] = DayItem{dayItem.ID, dayItem.Day, dayItem.ProjectID,  dayItem.ProjectItemID, "", dayItem.Status, dayItem.Text, dayItem.Sort}
	}
	return dayItems, nil
}

func (interactor *DayItemInteractor) Save(dayItem DayItem) (DayItem, error) {
	// validate 
	if dayItem.Text == "" {
		err := errors.New("Text is required")
		return dayItem, err
	}
	
  // either save or create
	entity := domain.DayItem{}
  if dayItem.ID != "" {
    // get the current entity
    entity, _ = interactor.Context.DayItems.Get(dayItem.ID)
  } else {
    // setup the new record
    entity.Day = dayItem.Day
    entity.ProjectID = dayItem.ProjectID
    entity.ProjectItemID = dayItem.ProjectItemID
    entity.Start = time.Now()
    entity.End = time.Now()
    entity.Duration = 0
  }
	entity.Status = dayItem.Status
  entity.Text = dayItem.Text
  
	// save
	storedEntity, err := interactor.Context.DayItems.Store(entity)
	if err == nil {
		dayItem.ID = storedEntity.ID
	}
	
	return dayItem, err
}


