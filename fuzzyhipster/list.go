package fuzzyhipster

import (
    "log"
    "net/http"
    "encoding/json"
    "strconv"
    //"src/usecases"
	  "github.com/gorilla/mux"
)

type List struct {
  Id      string      `json:"id"`
  Title    string     `json:"title"`
  Description string  `json:"description"`
  //Items  []int        `json:"items"`
}

type ListJSON struct {
  List List `json:"list"`
}

type ListsJSON struct {
  Lists []List `json:"lists"`
}

type Item struct {
  Id      string    `json:"id"`
  ListId  string    `json:"listid"`
  Name    string    `json:"name"`
  IsDone  bool      `json:"isDone"`
}

type ItemJSON struct {
  Item Item `json:"item"`
}

type ItemsJSON struct {
  Items []Item `json:"items"`
}

var idCounter int
var lists []List
var items []Item

func InitList() {
  if lists == nil {
      lists = []List{}
    }
  

	list := List{
    Id: strconv.Itoa(idCounter),
    Title: "Title One",
    Description: "Description One",
  }
  
  idCounter++    

  lists = append(lists, list)
  
  lists = append(lists, List{
    Id: strconv.Itoa(idCounter),
    Title: "Title Two",
    Description: "Description Two",
  })
  
  newLists := []List{
    {"a", "A", "AA", },
    {"b", "B", "BB", },
  }
  
  lists = append(lists, newLists...)
  
  idCounter++ 
  
}

func ListHandler(w http.ResponseWriter, r *http.Request) {
      
    w.Header().Set("Content-Type", "application/json")
  
    vars := mux.Vars(r)
    id := vars["id"]

	  // Find the index of the list
	  listIndex := -1
	  for index, _ := range lists {
      if lists[index].Id == id {
			  listIndex = index
			  break
		  }
	  }

	  // If we actually found a list remove it from the slice
    
    var listItem List
	  if listIndex != -1 {
      listItem = lists[listIndex]
	  }
    j, err := json.Marshal(ListJSON{List: listItem})
    if err != nil {
      panic(err)
    }
    w.Write(j)
}

func ListsHandler(w http.ResponseWriter, r *http.Request) {
    if lists == nil {
      lists = []List{}
    }  
    w.Header().Set("Content-Type", "application/json")
    j, err := json.Marshal(ListsJSON{Lists: lists})
    if err != nil {
      panic(err)
    }
    w.Write(j)
}

func CreateListHandler(w http.ResponseWriter, r *http.Request) {
  log.Println(r.Body)
  
	var listJSON ListJSON
	err := json.NewDecoder(r.Body).Decode(&listJSON)
	if err != nil {
		panic(err)
	}

	idCounter++

	list := listJSON.List
  list.Id = strconv.Itoa(idCounter)
  
	lists = append(lists, list)

	// Serialize the modified kitten to JSON
	j, err := json.Marshal(ListJSON{List: list})
	if err != nil {
		panic(err)
	}

  log.Println(j)
  
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func UpdateListHandler(w http.ResponseWriter, r *http.Request) {
  
	// Grab the kitten's id from the incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Decode the incoming list json
	var listJSON ListJSON
  err := json.NewDecoder(r.Body).Decode(&listJSON)
	if err != nil {
		panic(err)
	}

	// Find the kitten in our kittens slice and upate it's name
	for index, _ := range lists {
		if lists[index].Id == id {
			lists[index].Title = listJSON.List.Title
      lists[index].Description = listJSON.List.Description
		}
	}

	// Respond with a 204 indicating success, but no content
	w.WriteHeader(http.StatusNoContent)
}

func DeleteListHandler(w http.ResponseWriter, r *http.Request) {
	// Grab the list's id from the incoming url
	vars := mux.Vars(r)
	//id, err := strconv.Atoi(vars["id"])
  id := vars["id"]
	//if err != nil {
	//	panic(err)
	//}

	// Find the index of the list
	listIndex := -1
	for index, _ := range lists {
    //idToCompare, _ := strconv.Atoi(index)
		if lists[index].Id == id {
			listIndex = index
			break
		}
	}

	// If we actually found a list remove it from the slice
	if listIndex != -1 {
		lists = append(lists[:listIndex], lists[listIndex+1:]...)
	}

	// Respond with a 204 indicating success, but no content
	w.WriteHeader(http.StatusNoContent)
}

func ItemsHandler(w http.ResponseWriter, r *http.Request) {
    if items == nil {
      items = []Item{}
    }  
    w.Header().Set("Content-Type", "application/json")
    j, err := json.Marshal(ItemsJSON{Items: items})
    if err != nil {
      panic(err)
    }
    w.Write(j)
}

func ItemHandler(w http.ResponseWriter, r *http.Request) {
    if items == nil {
      items = []Item{}
    }  
    w.Header().Set("Content-Type", "application/json")
  
    vars := mux.Vars(r)
    id := vars["id"]

	  // Find the index of the list
	  itemIndex := -1
	  for index, _ := range items {
      if items[index].Id == id {
			  itemIndex = index
			  break
		  }
	  }

	  // If we actually found a list remove it from the slice
    
    var itemItem Item
	  if itemIndex != -1 {
      itemItem = items[itemIndex]
	  }
    j, err := json.Marshal(ItemJSON{Item: itemItem})
    if err != nil {
      panic(err)
    }
    w.Write(j)
}

func CreateItemHandler(w http.ResponseWriter, r *http.Request) {
	var listJSON ListJSON
	err := json.NewDecoder(r.Body).Decode(&listJSON)
	if err != nil {
		panic(err)
	}

	idCounter++

	list := listJSON.List
  list.Id = strconv.Itoa(idCounter)
  
	lists = append(lists, list)

	// Serialize the modified kitten to JSON
	j, err := json.Marshal(ListJSON{List: list})
	if err != nil {
		panic(err)
	}

  log.Println(j)
  
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
