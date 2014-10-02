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
  "vendor/github.com/gorilla/mux"
  //"vendor/github.com/dgrijalva/jwt-go"
)

type ProjectViewModel struct {
  Project usecases.Project
  Items []usecases.ProjectLine
}

func projectslistHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {  
	log.Println("projects:get")
	projects, err1 := u.Projects.FindActive()
  	if err1 != nil {
    	log.Println(err1)
  	}
  	if projects == nil {
      	projects = []usecases.Project{}
	} 
  
  // build the page view model
  page := buildPage(r, u)
  page.Model = projects
  page.IsProjectView = true
  
  render(w, "projects", page)  
}

func projectHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {  
	vars := mux.Vars(r)
	id := vars["id"]
  	// get the project
  	project, errFind := u.Projects.FindByID(id)
  	if errFind != nil {
    	http.Redirect(w, r, "/projects", http.StatusFound)
    	return
  	}
    viewModel := &ProjectViewModel{Project: project}
  	render(w, "project", &Page{Title: "Project", IsProjectView: true, Model: viewModel })
}

func projectAddHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {  
	log.Println("projectadd:get")
	render(w, "projectadd", &Page{Title: "Add Project", IsProjectView: true, Model: usecases.Project{} })
}

func projectAddPostHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {  
	// add an error
	log.Println("projectadd:post")
	
	errForm := r.ParseForm()
	if errForm != nil {
		render(w, "projectadd", &Page{Title: "Add Project", IsProjectView: true, Error: errForm.Error(), Model: &usecases.Project{} })
  		return
	}
	
	//r.Form.Get("")
	// get the lines that were submitted at the same time
		
	project := usecases.Project{ Title: r.Form.Get("Title") } //r.FormValue("Title") } 
  	createdProject, err := u.Projects.Save(project)
  	if err != nil {
    	render(w, "projectadd", &Page{Title: "Add Project", IsProjectView: true, Error: err.Error(), Model: &usecases.Project{} })
  		return
  	}
  	http.Redirect(w, r, "/project/" + createdProject.ID, http.StatusFound)
}

func projectPostHandler() {

}