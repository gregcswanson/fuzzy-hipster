package usecases

import (
	"src/interfaces"
  //"src/domain"
  //"time"
  //"errors"
  "log"
  //"sort"
)

type DaySummary struct {
	ID string
	Day int
	Text string
}

type DaySummaryInteractor struct {
	Context interfaces.DomainContext
}

func (interactor *DaySummaryInteractor) FindByDay(day int) ([]DaySummary, error) {
	// get the items on the month
  log.Println("usecases.DaySummaryInteractor.FindByDay")
  //daySummary, findError := interactor.Context.DaySummaries.FindForDay(day)
  //if findError != nil {
  //  log.Println(findError)
  //  return []DaySummary{}, findError
  //}
  
	// Copy to the use case model
  //var monthItems []MonthItem
  //sort.Sort(domain.MonthItemBySort(foundMonthItems))
	//monthItems = make([]MonthItem, len(foundMonthItems))
	//for i, monthItem := range foundMonthItems {
    //monthItems[i] = MonthItem{monthItem.ID, monthItem.Year, monthItem.Month, monthItem.ProjectID,  monthItem.ProjectItemID, "", monthItem.Status, monthItem.Text, monthItem.Sort}
	//}
	//return monthItems, nil
	return []DaySummary{}, nil
}

func (interactor *DaySummaryInteractor) Save(daySummary DaySummary) (DaySummary, error) {
  
  // update the month and month summary
  /*
	// validate 
	if monthItem.Text == "" {
		err := errors.New("Text is required")
		return monthItem, err
	}
  
  allMonthItems, _ := interactor.Context.MonthItems.Find(monthItem.Year, monthItem.Month)
	
  // either save or create
  if monthItem.ID != "" {
    log.Println("Usecase update month item", monthItem.Text)
    // get the current entity
    i, findError := domain.MonthItems(allMonthItems).Find(monthItem.ID)
    if findError != nil {
      log.Println(findError)
      return monthItem, findError
    }
    if(allMonthItems[i].Sort < monthItem.Sort ) {
      allMonthItems[i].Sort = monthItem.Sort + 5
    }else if (allMonthItems[i].Sort > monthItem.Sort) {
      allMonthItems[i].Sort = monthItem.Sort - 5
    }    
	  allMonthItems[i].Status = monthItem.Status 
    allMonthItems[i].Text = monthItem.Text
  } else {
    // setup the new record
    log.Println("Usecase insert item", monthItem.Text)
    entity := domain.MonthItem{}
    entity.Year = monthItem.Year
    entity.Month = monthItem.Month
    entity.ProjectID = monthItem.ProjectID
    entity.ProjectItemID = monthItem.ProjectItemID
    entity.Start = time.Now()
    entity.End = time.Now()
    entity.Duration = 0
	  entity.Status = monthItem.Status
    entity.Text = monthItem.Text
    entity.Sort = time.Now().Unix()
    allMonthItems = append(allMonthItems, entity)
  }
  
  sort.Sort(domain.MonthItemBySort(allMonthItems))
    var position int64 
    for _, otherItem := range allMonthItems {
      position = position + 10
      log.Println(otherItem.Sort, position, otherItem.Status, otherItem.Text)
      otherItem.Sort = position
      storedEntity, err := interactor.Context.MonthItems.Store(otherItem)
      if err == nil {
        if(monthItem.ID == ""){
          monthItem.ID = storedEntity.ID
        }
	    } else {
        log.Println("ERROR",err.Error())
      }
	  }
	*/
	return daySummary, nil
}



