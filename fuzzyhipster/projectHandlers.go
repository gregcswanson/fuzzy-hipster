package fuzzyhipster

import (
    "log"
    "net/http"
    "encoding/json"
    //"strconv"
    "src/usecases"
 // "fmt"
	//"github.com/gorilla/mux"
)

type ProjectJSON struct {
  Project usecases.Project `json:"list"`
}

type ProjectsJSON struct {
	Projects []usecases.Project `json:"lists"`
}

func ProjectsHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {  
  w.Header().Set("Content-Type", "application/json")
  log.Println("ProjectsHandler.FindActive")
  projects, _ := u.Projects.FindActive()
  log.Println("ProjectsHandler.Active")
  j, err := json.Marshal(ProjectsJSON{Projects: projects})
  if err != nil {
    panic(err)
  }
  w.Write(j)
}

