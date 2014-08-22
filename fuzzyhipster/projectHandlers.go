package fuzzyhipster

import (
    "log"
    "net/http"
    "encoding/json"
    //"strconv"
    "src/usecases"
  "fmt"
	"github.com/gorilla/mux"
  "errors"
)

type ProjectJSON struct {
  Project usecases.Project `json:"project"`
}

type ProjectLineJSON struct {
  ProjectLine usecases.ProjectLine `json:"line"`
}

type ProjectsJSON struct {
	Projects []usecases.Project `json:"projects"`
}

func ProjectsHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {  
  w.Header().Set("Content-Type", "application/json")
  log.Println("ProjectsHandler.FindActive")
  projects, err1 := u.Projects.FindActive()
  if err1 != nil {
    log.Println(err1)
  }
  if projects == nil {
      projects = []usecases.Project{}
    } 
  j, err := json.Marshal(ProjectsJSON{Projects: projects})
  if err != nil {
    panic(err)
  }
  w.Write(j)
}

func ProjectHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {
  w.Header().Set("Content-Type", "application/json")
  log.Println("ProjectHandler")
  // get the id
  vars := mux.Vars(r)
  id := vars["id"]
  // get the project
  project, errFind := u.Projects.FindByID(id)
  if errFind != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprint(w, errFind)
    return
  }
  j, err := json.Marshal(ProjectJSON{Project: project})
  if err != nil {
    panic(err)
  }
  w.Write(j)
}

func CreateProjectHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {
	var projectJSON ProjectJSON
  log.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&projectJSON)
	if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprint(w, err)
    return
	}
  
  log.Println(projectJSON.Project.Title)

	project := projectJSON.Project
  createProject, err := u.Projects.Save(project)
  if err != nil {
    log.Println(err)
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprint(w, err)
    return
  }
  
  log.Println(createProject.ID)
  
	// Serialize the modified project to JSON
	j, err := json.Marshal(ProjectJSON{Project: project})
	if err != nil {
		panic(err)
	}

  log.Println(j)
  
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func CreateProjectLineHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {
	var projectLineJSON ProjectLineJSON
  
  vars := mux.Vars(r)
  projectId := vars["project_id"]
  
	err := json.NewDecoder(r.Body).Decode(&projectLineJSON)
	if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprint(w, err)
    return
	}
  
	projectLine := projectLineJSON.ProjectLine
  createdLine, err := u.Projects.SaveItem(projectId, projectLine)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprint(w, err)
    return
  }
  
  log.Println(createdLine.ID)
  
	// Serialize the modified project to JSON
	j, err := json.Marshal(ProjectLineJSON{ProjectLine: createdLine})
	if err != nil {
		panic(err)
	}
  
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func UpdateProjectLineHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {
	var projectLineJSON ProjectLineJSON
  
  log.Println("UpdateProjectLineHandler")
  
  vars := mux.Vars(r)
  projectId := vars["project_id"]
  id := vars["id"]
  
	err := json.NewDecoder(r.Body).Decode(&projectLineJSON)
	if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprint(w, err)
    return
	}
  
	projectLine := projectLineJSON.ProjectLine
  
  if projectLine.ID != id {
    err := errors.New("Fail")
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprint(w, err)
    return
  }
    
  updatedLine, err := u.Projects.SaveItem(projectId, projectLine)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprint(w, err)
    return
  }
    
	// Serialize the modified project to JSON
	j, err := json.Marshal(ProjectLineJSON{ProjectLine: updatedLine})
	if err != nil {
		panic(err)
	}
  
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func DeleteProjectLineHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {
	var projectLineJSON ProjectLineJSON
  
  log.Println("DeleteProjectLineHandler")
  
  vars := mux.Vars(r)
  id := vars["id"]
  
	err := json.NewDecoder(r.Body).Decode(&projectLineJSON)
	if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprint(w, err)
    return
	}
  
	projectLine := projectLineJSON.ProjectLine
  
  if projectLine.ID != id {
    err := errors.New("Fail")
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprint(w, err)
    return
  }
    
  err = u.Projects.DeleteItem(projectLine.ID)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprint(w, err)
    return
  }
    
	// Serialize the original project item to JSON and return it
	j, err := json.Marshal(ProjectLineJSON{ProjectLine: projectLine})
	if err != nil {
		panic(err)
	}
  
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}