package fuzzyhipster

import (
  //"fmt"
  "net/http"
  "src/usecases"
  "time"
  "log"
  //"errors"
  "strconv"
  //"encoding/json"
  //"vendor/github.com/gorilla/mux"
  //"vendor/github.com/dgrijalva/jwt-go"
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

	render(w, "index", &Page{Title: "Index", IsDayView: true, Model: indexPage })
}