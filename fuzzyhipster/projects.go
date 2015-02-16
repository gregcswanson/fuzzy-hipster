package fuzzyhipster

import (
  "net/http"
  "src/usecases"
  "time"
  "log"
  "strings"
  "strconv"
  "vendor/github.com/gorilla/mux"
)

type ProjectViewModel struct {
  Project usecases.Project
  Items []usecases.ProjectLine
}

type ProjectItemPage struct {
  Project usecases.Project
	Item usecases.ProjectLine
  Sort []ProjectSort
}

type ProjectSort struct {
  Position int
  Sort int64
  Text string
  Selected bool
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

func projectPostHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {
  // add an error
	vars := mux.Vars(r)
	id := vars["id"]
  
	errForm := r.ParseForm()
	if errForm != nil {
    flashError(r, u.User.Current().Id , errForm.Error())
		http.Redirect(w, r, "/", http.StatusFound)
  	return
	}
  text := r.Form.Get("newItemText")
  status := r.Form.Get("newItemType")
  if len(status) == 0 {
    status = "OPEN"
  }
  if strings.HasPrefix(text, "/") {
    text = strings.TrimPrefix(text, "/")
    status = "NOTE"
  } else if strings.HasPrefix(text, ".") {
    text = strings.TrimPrefix(text, ".")
    status = "OPEN"
  } else if strings.HasPrefix(text, "*") {
    text = strings.TrimPrefix(text, "*")
    status = "LABEL"
  }
  
  log.Println("new line item:", text)
			
  projectItem := usecases.ProjectLine{ ProjectID: id, Text: text, Sort: 0, Status: status } 
  _, err := u.Projects.SaveItem2(projectItem)
  if err != nil {
    flashError(r, u.User.Current().Id ,err.Error())
  } 
    
  http.Redirect(w, r, "/project/" + id, http.StatusFound)
}

func projectItemHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {
  // get the current day
  vars := mux.Vars(r)
	itemId := vars["item_id"]
  projectId := vars["project_id"]
  
	// get the project and line
	project, errProject := u.Projects.FindByID(projectId)
  if errProject != nil {
    log.Println(errProject)
  }
  editPage := ProjectItemPage{ Project: project }
  
  projectLineIndex, errLine := usecases.ProjectLines(project.Lines).Find(itemId)
  if errLine != nil {
    log.Println(errLine)
  } else  {
    editPage.Item = project.Lines[projectLineIndex]
  }
  
  // create the sort order
  editPage.Sort = make([]ProjectSort, len(project.Lines))
	for i, items := range project.Lines {
    editPage.Sort[i] = ProjectSort{i + 1, items.Sort, items.Text, false}
    if items.ID == editPage.Item.ID {
      editPage.Sort[i].Selected = true
      log.Println("Sort Selected ", items.Sort)
    }
	}
  
  // setup the master page
  page := buildPage(r, u, time.Now())
  page.Title = "Project"
  page.Model = editPage
  page.IsProjectsView = true
  page.IsProjectView = true
  
	render(w, "projectitemedit", page)
}

func projectItemPostHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {  
	// add an error
	vars := mux.Vars(r)
	projectId := vars["project_id"]
  itemId := vars["item_id"]
  
	errForm := r.ParseForm()
	if errForm != nil {
    flashError(r, u.User.Current().Id , errForm.Error())
		http.Redirect(w, r, "/project/" + projectId, http.StatusFound)
  	return
	}
  
  line, err := u.Projects.FindItemById(itemId)
  if err != nil {
    flashError(r, u.User.Current().Id, err.Error())
		http.Redirect(w, r, "/project/" + projectId, http.StatusFound)
    return
  }
  
  line.Text = r.Form.Get("Text")
  line.Status = r.Form.Get("Status")
  sort, _ := strconv.Atoi(r.Form.Get("Sort"))
  line.Sort = int64(sort)
  
  _, errSave := u.Projects.SaveItem2(line)
  if errSave != nil {
    flashError(r, u.User.Current().Id, err.Error())
  }
    
  http.Redirect(w, r, "/project/" + projectId, http.StatusFound)
}

func projectItemTogglePostHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {  
	// add an error
	vars := mux.Vars(r)
	projectId := vars["project_id"]
  itemId := vars["item_id"]
  
  err := u.Projects.Toggle(itemId)
  if err != nil {
    flashError(r, u.User.Current().Id, err.Error())
  }
    
  http.Redirect(w, r, "/project/" + projectId, http.StatusFound)
}

func projectItemDeleteHandler(w http.ResponseWriter, r *http.Request, u *usecases.Interactors) {  
	// add an error
	vars := mux.Vars(r)
	projectId := vars["project_id"]
  itemId := vars["item_id"]
  
  err := u.Projects.DeleteItem(itemId)
  if err != nil {
    flashError(r, u.User.Current().Id, err.Error())
  }
    
  http.Redirect(w, r, "/project/" + projectId, http.StatusFound)
}
