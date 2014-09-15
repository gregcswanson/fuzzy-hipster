package fuzzyhipster

import (
  //"fmt"
  "net/http"
  "src/usecases"
  //"time"
  "log"
  //"errors"
  //"strconv"
  //"encoding/json"
  //"vendor/github.com/gorilla/mux"
  //"vendor/github.com/dgrijalva/jwt-go"
)

func projectslistHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {  
	
	projects, err1 := u.Projects.FindActive()
  	if err1 != nil {
    	log.Println(err1)
  	}
  	if projects == nil {
      	projects = []usecases.Project{}
	} 

	render(w, "projects", &Page{Title: "Index", IsProjectView: true, Model: projects })
}

func projectAddHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {  
	log.Println("projectadd:get")
	render(w, "projectadd", &Page{Title: "Add Project", IsProjectView: true, Model: usecases.Project{} })
}

func projectAddPostHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {  
	// add an error
	project := usecases.Project{ Title: r.FormValue("Title") } 
  	_, err := u.Projects.Save(project)
  	if err != nil {
    	render(w, "projectadd", &Page{Title: "Add Project", IsProjectView: true, Error: err.Error(), Model: &usecases.Project{} })
  		return
  	}
	w.Header().Set("Location", "/projects")
	w.WriteHeader(http.StatusFound)
}