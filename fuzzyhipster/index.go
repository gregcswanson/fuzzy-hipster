package fuzzyhipster

import (
  //"fmt"
  "net/http"
  "src/usecases"
  "time"
  "log"
  "strconv"
  "strings"
)

type IndexPage struct {
	DateAsInt int
	DateDisplay string
	DayItems []usecases.DayItem
}

func indexHander(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {  
	// get the current day
	dateAsInt, _ := strconv.Atoi(time.Now().Format("20060102"))
	indexPage := &IndexPage{ DateAsInt: dateAsInt, DateDisplay: time.Now().Format("Mon Jan 1, 2006") }

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

func indexPostHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {  
	// add an error
	log.Println("indexPostHandler")
	
  dateAsInt, _ := strconv.Atoi(time.Now().Format("20060102"))
  
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
    
  	http.Redirect(w, r, "/", http.StatusFound)
}