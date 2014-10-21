package usecases

import (
	"src/interfaces"
  "src/domain"
  "time"
  "errors"
  "log"
  "strings"
  "strconv"
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

type MonthDay struct {
  DateAsString string
  Day int
  DayCode string
  Display string
  Name string
  ShortName string
  Selected bool
  HasItems bool
  HasOpenItems bool
  Summary string
  Placeholder string
}

type Month struct {
  Selected bool
  Display string
  Name string
  ShortName string
  MonthCode string
  SelectedDay MonthDay
  Days []MonthDay
  HasItems bool
  HasOpenItems bool
}

type Year struct {
  Year int
  Selected Month
  Months []Month
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

func (interactor *DayItemInteractor) FindById(itemId string) (DayItem, error) {
  // get the item
  dayItem, err := interactor.Context.DayItems.Get(itemId)
  if err != nil {
    return DayItem{}, err
  }
  
  day := DayItem{dayItem.ID, dayItem.Day, dayItem.ProjectID,  
                 dayItem.ProjectItemID, "", dayItem.Status, dayItem.Text, dayItem.Sort}
  return day, nil

}

func (interactor *DayItemInteractor) Toggle(itemId string) error {
  log.Println(itemId)
  // get the item
  dayItem, err := interactor.Context.DayItems.Get(itemId)
  if err != nil {
    return err
  }
  if dayItem.Status == "OPEN" {
    dayItem.Status = "DONE"
  } else if dayItem.Status == "DONE" {
    dayItem.Status = "CANCELLED"
  } else if dayItem.Status == "CANCELLED" {
    dayItem.Status = "OPEN"
  }
  _ , errSave := interactor.Context.DayItems.Store(dayItem)
  return  errSave
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

func (interactor *DayItemInteractor) Delete(itemId string) error {
  // get the item
  errSave := interactor.Context.DayItems.Delete(itemId)
  return  errSave
}

func (interactor *DayItemInteractor) FindMonth(date time.Time) (Month, error) {
  // initialise
  month := Month{ Days: []MonthDay{} }
  
  monthItems, _ := interactor.Context.DayItems.FindMonth(date.Year(), int(date.Month()))
  log.Println(len(monthItems))
  
  for d := time.Date(date.Year(), date.Month(), 1, 23, 0, 0, 0, time.UTC); d.Month() == date.Month(); d = d.AddDate(0,0,1) {
    dayCode := string(d.Format("Mon")[0])
    dayNumber := d.Format("2")
    display := strings.Join([]string{dayNumber, dayCode}, " ")
    o := ordinal(d.Day())
    monthDay := MonthDay{ 
      DateAsString: d.Format("20060102"), 
      Day: d.Day(), 
      Selected: d.Format("20060102") == date.Format("20060102"), 
      Display: display, 
      Name: d.Format("Monday, 2" + o ),
      ShortName: d.Format("Mon, 2"),
      DayCode: dayCode }
    
    if monthDay.Selected {
      month.SelectedDay = monthDay
    }
    
    // to do - get the has items and open items details
    for _, value := range monthItems {
      dayString := strconv.Itoa(value.Day)
      compareString := d.Format("20060102")
      if dayString == compareString {
        monthDay.HasItems = true
        if value.Status == "OPEN" {
          monthDay.HasOpenItems = true
        }
      }
    }
    // get all items for the month before this method so the data is only hit once
    month.Days = append(month.Days, monthDay)
  }
  
  return month, nil
}

func (interactor *DayItemInteractor) FindYear(date time.Time) (Year, error) {
  // return the year summary
  year := Year{ Months: []Month{}, Year: date.Year() }
  
  // loop through 12 months to create the structure
  for i := 1; i < 13; i++ {
    d := time.Date(date.Year(), time.Month(i), 1, 23, 0, 0, 0, time.UTC);
    month := Month{ 
      MonthCode: d.Format("200601"), 
      Display: d.Format("Jan"), 
      Name: d.Format("January"),
      ShortName: d.Format("Jan"),
    }
    if d.Month() == date.Month() {
      month.Selected = true
      year.Selected = month
    }
    year.Months = append(year.Months, month)
  }
  
  // just a flag with is open or has items for the entire month  
  return year, nil
}

func ordinal(x int) string {
  suffix := "th"
  switch x % 10 {
    case 1: if (x % 100 != 11) { suffix = "st" }
    case 2: if (x % 100 != 12) { suffix = "nd" }
    case 3: if (x % 100 != 13) { suffix = "rd" }
  }
  return suffix
}


