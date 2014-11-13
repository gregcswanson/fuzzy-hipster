package fuzzyhipster

import (
  //"fmt"
  "net/http"
  "src/usecases"
  "time"
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
  page := buildPage(r, u, time.Now())
  page.Model = projects
  page.IsProjectsView = true
  
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
  	render(w, "project", &Page{Title: "Project", IsProjectsView: true, IsProjectView: true, Model: viewModel })
}

func projectUpsertHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {  
  // either add or edit the project depending on if the id parameter is included
  vars := mux.Vars(r)
	id := vars["id"]
  model := usecases.Project{}
  if id != "" {
    model, _ = u.Projects.FindByID(id)
  }
	log.Println("projectadd:upsert")
	render(w, "projectadd", &Page{Title: "Project", IsProjectsView: true, Model: model })
}

func projectUpsertPostHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {  
	// add an error
	log.Println("projectadd:post")
	
	errForm := r.ParseForm()
	if errForm != nil {
		render(w, "projectadd", &Page{Title: "Add Project", IsProjectsView: true, IsProjectView: true, Error: errForm.Error(), Model: &usecases.Project{} })
  		return
	}
	
	// get the lines that were submitted at the same time
		
  project := usecases.Project{ ID: r.Form.Get("ID"), Title: r.Form.Get("Title"), Description: r.Form.Get("Description") } 
  	createdProject, err := u.Projects.Save(project)
  	if err != nil {
    	render(w, "projectadd", &Page{Title: "Project", IsProjectsView: true, IsProjectView: true, Error: err.Error(), Model: &usecases.Project{} })
  		return
  	}
  	http.Redirect(w, r, "/project/" + createdProject.ID, http.StatusFound)
}

func projectDeleteHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {  
  // move the project to trash
  vars := mux.Vars(r)
	id := vars["id"]
  project, _ := u.Projects.FindByID(id)
  
  deleteError := u.Projects.Delete(project.ID)
  if deleteError != nil {
    flashError(r, u.User.Current().Id, deleteError.Error())
  }
  
  // set the message with undo - make a specific message type that has a link to trash
  flashWarning(r, u.User.Current().Id, "Project '" + project.Title + "' was moved to trash")
  
  // redirect to projects
  http.Redirect(w, r, "/projects", http.StatusFound)
}

func projectPostHandler() {

}