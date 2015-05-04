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

type MonthPage struct {
  MonthID string
	DateAsInt int
	DateDisplay string
  Month usecases.Month
  Year usecases.Year
	MonthItems []usecases.MonthItem
}

type MonthItemPage struct {
  MonthID string
	Item usecases.MonthItem
  Sort []Sort
}

func monthHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) { 
   vars := mux.Vars(r)
	id := vars["month_id"] + "01"
  
  selectedDate, _ := time.Parse("20060102", id)
  dateAsInt, errDay := strconv.Atoi(id)
  if errDay != nil {
    http.Redirect(w, r, "/", http.StatusFound)
    return
  }
  monthPage := &MonthPage{ DateAsInt: dateAsInt, DateDisplay: id } 
    
  monthItems, errMonthItems := u.MonthItems.FindByMonth(selectedDate.Year(), int(selectedDate.Month()))
  if errMonthItems != nil {
    log.Println(errMonthItems)
  }
  if monthItems == nil {
    monthItems = []usecases.MonthItem{}
  } 
  monthPage.MonthItems = monthItems
  monthPage.MonthID = vars["month_id"]
  monthPage.Month, _ = u.DayItems.FindMonth(selectedDate)
  monthPage.Year, _ = u.DayItems.FindYear(selectedDate)
  
  page := buildPage(r, u, selectedDate)
  page.Title = "Month"
  page.Model = monthPage
  page.IsMonthView = true
  
  render(w, "monthindex", page)  
}


func monthOverviewHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {  
	// get the current day
  vars := mux.Vars(r)
	id := vars["month_id"] + "01"
  
  selectedDate, _ := time.Parse("20060102", id)
  dateAsInt, errDay := strconv.Atoi(id)
  if errDay != nil {
    http.Redirect(w, r, "/", http.StatusFound)
    return
  }
  monthPage := &MonthPage{ DateAsInt: dateAsInt, DateDisplay: id } 

  // get the month summary
  monthPage.MonthID = vars["month_id"]
  monthPage.Month, _ = u.DayItems.FindMonth(selectedDate)
  monthPage.Year, _ = u.DayItems.FindYear(selectedDate)
  
  // setup the master page
  page := buildPage(r, u, selectedDate)
  page.Title = "Month"
  page.Model = monthPage
  page.IsMonthView = true
  page.IsMonthOverView = true
  
	render(w, "monthoverview", page)  
}

func monthItemsHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) { 
   vars := mux.Vars(r)
	id := vars["month_id"] + "01"
  
  selectedDate, _ := time.Parse("20060102", id)
  dateAsInt, errDay := strconv.Atoi(id)
  if errDay != nil {
    http.Redirect(w, r, "/", http.StatusFound)
    return
  }
  monthPage := &MonthPage{ DateAsInt: dateAsInt, DateDisplay: id } 
    
  monthItems, errMonthItems := u.MonthItems.FindByMonth(selectedDate.Year(), int(selectedDate.Month()))
  if errMonthItems != nil {
    log.Println(errMonthItems)
  }
  if monthItems == nil {
    monthItems = []usecases.MonthItem{}
  } 
  monthPage.MonthItems = monthItems
  monthPage.MonthID = vars["month_id"]
  
  page := buildPage(r, u, selectedDate)
  page.Title = "Month"
  page.Model = monthPage
  page.IsMonthView = true
  
  render(w, "monthitems", page)  
}

func monthItemsPostHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {  
	// add an error
	vars := mux.Vars(r)
	id := vars["month_id"]
  
  selectedDate, _ := time.Parse("20060102", id + "01")
  
  //dateAsInt, errDay := strconv.Atoi(id) //time.Now().Format("20060102"))
  //if errDay != nil {
  //  http.Redirect(w, r, "/", http.StatusFound)
  //  return
  //}
  
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
  		
  monthItem := usecases.MonthItem{ Month: int(selectedDate.Month()), Year: selectedDate.Year(), Text: text, Sort: 0, Status: status } 
  	_, err := u.MonthItems.Save(monthItem)
  	if err != nil {
    	flashError(r, u.User.Current().Id ,err.Error())
  	}
    
  	http.Redirect(w, r, "/month/" + id + "/items", http.StatusFound)
}

func monthItemTogglePostHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {  
	// add an error
	vars := mux.Vars(r)
	id := vars["month_id"]
  itemId := vars["item_id"]
  
  err := u.MonthItems.Toggle(itemId)
  if err != nil {
    flashError(r, u.User.Current().Id, err.Error())
  }
    
  http.Redirect(w, r, "/month/" + id + "/items", http.StatusFound)
}

func monthItemHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {
  // get the current day
  vars := mux.Vars(r)
	id := vars["month_id"]
	itemId := vars["item_id"]
  
	monthItem, err := u.MonthItems.FindById(itemId)
  if err != nil {
    log.Println(err)
  }
  
  selectedDate, _ := time.Parse("20060102", id + "01")
  
  // setup the page
  monthItemPage := MonthItemPage{ Item: monthItem}
  monthItemPage.MonthID = id
  otherItems, _ := u.MonthItems.FindByMonth(selectedDate.Year(), int(selectedDate.Month()))
  
  // create the sort order
  monthItemPage.Sort = make([]Sort, len(otherItems))
	for i, otherItem := range otherItems {
    monthItemPage.Sort[i] = Sort{i + 1, otherItem.Sort, otherItem.Text, false}
    if otherItem.ID == monthItem.ID {
      monthItemPage.Sort[i].Selected = true
    }
	}
  
  // setup the master page
  page := buildPage(r, u, time.Now())
  page.Title = "Month Item"
  page.Model = monthItemPage
  page.IsMonthView = true
  
	render(w, "monthedit", page)
}

func monthItemPostHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {  
	// add an error
	vars := mux.Vars(r)
	monthId := vars["month_id"]
  itemId := vars["item_id"]
  
	errForm := r.ParseForm()
	if errForm != nil {
    flashError(r, u.User.Current().Id , errForm.Error())
		http.Redirect(w, r, "/", http.StatusFound)
  	return
	}
  
  monthItem, err := u.MonthItems.FindById(itemId)
  if err != nil {
    flashError(r, u.User.Current().Id, err.Error())
		http.Redirect(w, r, "/", http.StatusFound)
    return
  }
  
  monthItem.Text = r.Form.Get("Text")
  monthItem.Status = r.Form.Get("Status")
  sort, _ := strconv.Atoi(r.Form.Get("Sort"))
  monthItem.Sort = int64(sort)
  
  _, errSave := u.MonthItems.Save(monthItem)
  if errSave != nil {
    flashError(r, u.User.Current().Id, err.Error())
  }
    
  http.Redirect(w, r, "/month/" + monthId + "/items", http.StatusFound)
}

func monthItemDeleteHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {  
	// add an error
	vars := mux.Vars(r)
	monthId := vars["month_id"]
  itemId := vars["item_id"]
  
  err := u.MonthItems.Delete(itemId)
  if err != nil {
    flashError(r, u.User.Current().Id, err.Error())
  }
    
  http.Redirect(w, r, "/month/" + monthId + "/items", http.StatusFound)
}