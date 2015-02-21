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

type IndexPage struct {
	DateAsInt int
	DateDisplay string
  Year usecases.Year
  Month usecases.Month
	DayItems []usecases.DayItem
}

type EditPage struct {
	Item usecases.DayItem
  Sort []Sort
}

type Sort struct {
  Position int
  Sort int64
  Text string
  Selected bool
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
  
  selectedDate, _ := time.Parse("20060102", id)
  dateAsInt, errDay := strconv.Atoi(id)
  if errDay != nil {
    http.Redirect(w, r, "/", http.StatusFound)
    return
  }
  indexPage := &IndexPage{ DateAsInt: dateAsInt, DateDisplay: id } 

  // get the month summary
  indexPage.Month, _ = u.DayItems.FindMonth(selectedDate)
  indexPage.Year, _ = u.DayItems.FindYear(selectedDate)
  
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
  page := buildPage(r, u, selectedDate)
  page.Title = "Index"
  page.Model = indexPage
  page.IsDayView = true
  //page.IsMonthView = true
  
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
  } else if strings.HasPrefix(text, "*") {
    text = strings.TrimPrefix(text, "*")
    status = "LABEL"
  } else if strings.HasPrefix(text, "#") {
    text = strings.TrimPrefix(text, "#")
    status = "MEETING"
  }
			
  dayItem := usecases.DayItem{ Day: dateAsInt, Text: text, Sort: 0, Status: status } 
  	_, err := u.DayItems.Save(dayItem)
  	if err != nil {
    	flashError(r, u.User.Current().Id ,err.Error())
  	} //else {
      //flashMessage(r, u.User.Current().Id ,"item created")
    //}
    
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
  
  // setup the page
  editPage := EditPage{ Item: dayItem}
  otherItems, _ := u.DayItems.FindByDay(dayItem.Day)
  
  // create the sort order
  editPage.Sort = make([]Sort, len(otherItems))
	for i, otherItem := range otherItems {
    editPage.Sort[i] = Sort{i + 1, otherItem.Sort, otherItem.Text, false}
    if otherItem.ID == dayItem.ID {
      editPage.Sort[i].Selected = true
      log.Println("Sort Selected ", otherItem.Sort)
    }
	}
  
  // setup the master page
  page := buildPage(r, u, time.Now())
  page.Title = "Index"
  page.Model = editPage
  page.IsDayView = true
  page.IsMonthView = true
  
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
  
  log.Println(dayItem.Sort)
  dayItem.Text = r.Form.Get("Text")
  dayItem.Status = r.Form.Get("Status")
  log.Println(r.Form.Get("Status"))
  log.Println(r.Form.Get("Sort"))
  sort, _ := strconv.Atoi(r.Form.Get("Sort"))
  dayItem.Sort = int64(sort)
  log.Println(dayItem.Sort)
  
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
