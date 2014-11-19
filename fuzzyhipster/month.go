package fuzzyhipster

import (
  //"fmt"
  "net/http"
  "src/usecases"
  "time"
  //"log"
  "strconv"
  //"strings"
  "vendor/github.com/gorilla/mux"
)

type MonthPage struct {
	DateAsInt int
	DateDisplay string
  Month usecases.Month
  Year usecases.Year
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
    
  page := buildPage(r, u, selectedDate)
  page.Title = "Month"
  page.Model = monthPage
  page.IsMonthView = true
  
  render(w, "monthitems", page)  
}
