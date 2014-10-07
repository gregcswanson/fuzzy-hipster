package fuzzyhipster

import (
  //"fmt"
  "net/http"
  "src/usecases"
  "time"
  "log"
  "strconv"
  "strings"
  "vendor/github.com/gorilla/mux"
)

type NavigationDay struct {
  DateAsString string
  Day int
  DayCode string
  Display string
  Selected bool
  HasItems bool
  HasOpenItems bool
}

type Navigation struct {
  Days []NavigationDay
}

type IndexPage struct {
	DateAsInt int
	DateDisplay string
  Navigation Navigation
	DayItems []usecases.DayItem
}

func indexHander(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {  
	// get the current day
	dateAsString := time.Now().Format("20060102")
  http.Redirect(w, r, "/day/" + dateAsString , http.StatusFound)
}

func dayHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {  
	// get the current day
  
  vars := mux.Vars(r)
	id := vars["day_id"]
  
  dateAsInt, errDay := strconv.Atoi(id) //time.Now().Format("20060102"))
  if errDay != nil {
    http.Redirect(w, r, "/", http.StatusFound)
    return
  }
  indexPage := &IndexPage{ DateAsInt: dateAsInt, DateDisplay: id, Navigation: Navigation{} } 

  // REFACTOR INTO OWN FUNCTION THAT CREATED NAVIGATION
  // get all the day item in the current month to be used for summarising
  // get the days in the current month
  indexPage.Navigation.Days = []NavigationDay{}
  selectedDate, _ := time.Parse("20060102", id)
  for d := time.Date(selectedDate.Year(), selectedDate.Month(), 1, 23, 0, 0, 0, time.UTC); d.Month() == selectedDate.Month(); d = d.AddDate(0,0,1) {
    dayCode := string(d.Format("Mon")[0])
    dayNumber := d.Format("2")
    display := strings.Join([]string{dayNumber, dayCode}, " ")
    navigationDay := NavigationDay{ DateAsString: d.Format("20060102"), Day: d.Day(), Selected: d.Format("20060102") == id, Display: display, DayCode: dayCode }
    // to do - get the has items and open items details
    // get all items for the month before this method so the data is only hit once
    indexPage.Navigation.Days = append(indexPage.Navigation.Days, navigationDay)
  }
  
  
	// get the items for the current day
	dayItems, err1 := u.DayItems.FindByDay(dateAsInt)
  	if err1 != nil {
    	log.Println(err1)
  	}
  	if dayItems == nil {
      	dayItems = []usecases.DayItem{}
  	} 
  	indexPage.DayItems = dayItems
  
  // setup the master page
  page := buildPage(r, u)
  page.Title = "Index"
  page.Model = indexPage
  page.IsDayView = true
  
	render(w, "index", page)  
}

func dayPostHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {  
	// add an error
	vars := mux.Vars(r)
	id := vars["day_id"]
  
  dateAsInt, errDay := strconv.Atoi(id) //time.Now().Format("20060102"))
  if errDay != nil {
    http.Redirect(w, r, "/", http.StatusFound)
    return
  }
  
	errForm := r.ParseForm()
	if errForm != nil {
    flashError(r, u.User.Current().Id , errForm.Error())
		http.Redirect(w, r, "/", http.StatusFound)
  	return
	}
  text := r.Form.Get("newItemText")
  status := r.Form.Get("newItemType")
  if len(status) == 0 {
    status = "OPEN"
  }
  if strings.HasPrefix(text, "/") {
    text = strings.TrimPrefix(text, "/")
    status = "NOTE"
  } else if strings.HasPrefix(text, ".") {
    text = strings.TrimPrefix(text, ".")
    status = "OPEN"
  }
			
  dayItem := usecases.DayItem{ Day: dateAsInt, Text: text, Sort: 0, Status: status } 
  	_, err := u.DayItems.Save(dayItem)
  	if err != nil {
    	flashError(r, u.User.Current().Id ,err.Error())
  	} else {
      flashMessage(r, u.User.Current().Id ,"item created")
    }
    
  	http.Redirect(w, r, "/day/" + id, http.StatusFound)
}

func togglePostHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {  
	// add an error
	vars := mux.Vars(r)
	id := vars["day_id"]
  itemId := vars["item_id"]
  
  err := u.DayItems.Toggle(itemId)
  if err != nil {
    flashError(r, u.User.Current().Id, err.Error())
  }
    
  http.Redirect(w, r, "/day/" + id, http.StatusFound)
}

func dayItemHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {
  // get the current day
  vars := mux.Vars(r)
	itemId := vars["item_id"]
  
	// get the items for the current day
	dayItem, err := u.DayItems.FindById(itemId)
  if err != nil {
    log.Println(err)
  }
  
  // setup the master page
  page := buildPage(r, u)
  page.Title = "Index"
  page.Model = dayItem
  page.IsDayView = true
  
	render(w, "dayedit", page)
}

func dayItemPostHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {  
	// add an error
	vars := mux.Vars(r)
	dayId := vars["day_id"]
  itemId := vars["item_id"]
  
	errForm := r.ParseForm()
	if errForm != nil {
    flashError(r, u.User.Current().Id , errForm.Error())
		http.Redirect(w, r, "/", http.StatusFound)
  	return
	}
  
  dayItem, err := u.DayItems.FindById(itemId)
  if err != nil {
    flashError(r, u.User.Current().Id, err.Error())
		http.Redirect(w, r, "/", http.StatusFound)
    return
  }
  
  dayItem.Text = r.Form.Get("Text")
  dayItem.Status = r.Form.Get("Status")
  
  _, errSave := u.DayItems.Save(dayItem)
  if errSave != nil {
    flashError(r, u.User.Current().Id, err.Error())
  }
    
  http.Redirect(w, r, "/day/" + dayId, http.StatusFound)
}

func dayItemDeleteHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {  
	// add an error
	vars := mux.Vars(r)
	dayId := vars["day_id"]
  itemId := vars["item_id"]
  
  err := u.DayItems.Delete(itemId)
  if err != nil {
    flashError(r, u.User.Current().Id, err.Error())
  }
    
  http.Redirect(w, r, "/day/" + dayId, http.StatusFound)
}
