package fuzzyhipster

import (
  "log"
  "net/http"
  "encoding/json"
  "strconv"
  "src/usecases"
  "fmt"
	"github.com/gorilla/mux"
  //"errors"
)

type DayItemJSON struct {
  DayItem usecases.DayItem `json:"dayitem"`
}

type DayItemsJSON struct {
	DayItems []usecases.DayItem `json:"dayitems"`
}

func DayHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {  
  w.Header().Set("Content-Type", "application/json")
  log.Println("DayHandler.FindActive")
  
  vars := mux.Vars(r)
  id := vars["id"]
  
  // convert to int
  dayItemAsInt, errConversion := strconv.Atoi(id)
  if errConversion != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprint(w, "Day is invalid")
    return
  }
  
  dayItems, err1 := u.DayItems.FindByDay(dayItemAsInt)
  if err1 != nil {
    log.Println(err1)
  }
  if dayItems == nil {
      dayItems = []usecases.DayItem{}
  } 
  j, err := json.Marshal(DayItemsJSON{DayItems: dayItems})
  if err != nil {
    panic(err)
  }
  w.Write(j)
}

func CreateDayItemHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {
	var dayItemJSON DayItemJSON
  log.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&dayItemJSON)
	if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprint(w, err)
    return
	}
  
  log.Println(dayItemJSON.DayItem.Text)

	dayItem := dayItemJSON.DayItem
  createDayItem, err := u.DayItems.Save(dayItem)
  if err != nil {
    log.Println(err)
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprint(w, err)
    return
  }
  
  log.Println(createDayItem.ID)
  
	j, err := json.Marshal(DayItemJSON{DayItem: createDayItem})
	if err != nil {
		panic(err)
	}
  
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
