package fuzzyhipster

import (
    "log"
    "net/http"
    "encoding/json"
)

type CheckList struct {
    Id      int    `json:"id"`
    Name    string `json:"name"`
}

type CheckListJSON struct {
  CheckList CheckList `json:"checklist"`
}

type CheckListsJSON struct {
  CheckLists []CheckList `json:"checklists"`
}

var checklists []CheckList

func CheckListsHandler(w http.ResponseWriter, r *http.Request) {
    
    user, _ :=authenticateRequest(w, r);

    log.Println(user)
  
    w.Header().Set("Content-Type", "application/json")
    j, err := json.Marshal(CheckListsJSON{CheckLists: checklists})
    if err != nil {
      panic(err)
    }
    w.Write(j)
}