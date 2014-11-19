package fuzzyhipster

import (
  "net/http"
  "vendor/github.com/gorilla/mux"
)

func initRouter() {
  r := mux.NewRouter()
  // authentication
    
  r.HandleFunc("/api/1/gettoken", authenticated(handlerGetToken)).Methods("GET")
  r.HandleFunc("/api/1/checktoken", handlerCheckToken).Methods("GET")
  r.HandleFunc("/api/1/token", TokenHandler).Methods("GET")
  r.HandleFunc("/api/1/tokenread", TokenReadHander).Methods("GET")
  
  // projects
  r.HandleFunc("/api/1/projects", useCaseMiddleware(ProjectsHandler)).Methods("GET")
  r.HandleFunc("/api/1/projects/{id}", useCaseMiddleware(ProjectHandler)).Methods("GET")
  r.HandleFunc("/api/1/projects", useCaseMiddleware(CreateProjectHandler)).Methods("POST")
  r.HandleFunc("/api/1/projects/{id}", useCaseMiddleware(UpdateProjectHandler)).Methods("PUT")
  r.HandleFunc("/api/1/projects/{id}", DeleteListHandler).Methods("DELETE")
  
  // project lines
  r.HandleFunc("/api/1/projects/{project_id}/lines",  useCaseMiddleware(CreateProjectLineHandler)).Methods("POST")
  r.HandleFunc("/api/1/projects/{project_id}/lines/{id}",  useCaseMiddleware(UpdateProjectLineHandler)).Methods("PUT")
  r.HandleFunc("/api/1/projects/{project_id}/lines/{id}", useCaseMiddleware(DeleteProjectLineHandler)).Methods("DELETE")
  
  // day items
  r.HandleFunc("/api/1/day/{day_id}", useCaseMiddleware(DayHandler)).Methods("GET")
  //r.HandleFunc("/api/1/dayitem/{id}", useCaseMiddleware(ProjectHandler)).Methods("GET")
  r.HandleFunc("/api/1/dayitem", useCaseMiddleware(CreateDayItemHandler)).Methods("POST")
  //r.HandleFunc("/api/1/dayitem/{id}", useCaseMiddleware(UpdateProjectHandler)).Methods("PUT")
  //r.HandleFunc("/api/1/dayitem/{id}", DeleteListHandler).Methods("DELETE")
  
  // lists
  r.HandleFunc("/api/1/lists", ListsHandler).Methods("GET")
  r.HandleFunc("/api/1/lists/{id}", ListHandler).Methods("GET")
  r.HandleFunc("/api/1/lists", CreateListHandler).Methods("POST")
  r.HandleFunc("/api/1/lists/{id}", UpdateListHandler).Methods("PUT")
  r.HandleFunc("/api/1/lists/{id}", DeleteListHandler).Methods("DELETE")
  
  // items
  r.HandleFunc("/api/1/items", ItemsHandler).Methods("GET")
  r.HandleFunc("/api/1/items/{id}", ItemHandler).Methods("GET")
  
  r.HandleFunc("/", useCaseRequest(indexHander)).Methods("GET")
  r.HandleFunc("/day/{day_id}", useCaseRequest(dayHandler)).Methods("GET")
  r.HandleFunc("/day/{day_id}", useCaseRequest(dayPostHandler)).Methods("POST")
  r.HandleFunc("/day/{day_id}/item/{item_id}", useCaseRequest(dayItemHandler)).Methods("GET")
  r.HandleFunc("/day/{day_id}/item/{item_id}", useCaseRequest(dayItemPostHandler)).Methods("POST")
  r.HandleFunc("/day/{day_id}/item/{item_id}/toggle", useCaseRequest(togglePostHandler)).Methods("POST")
  r.HandleFunc("/day/{day_id}/item/{item_id}/delete", useCaseRequest(dayItemDeleteHandler)).Methods("GET")
  
  r.HandleFunc("/month/{month_id}/items", useCaseRequest(monthItemsHandler)).Methods("GET")
  r.HandleFunc("/month/{month_id}/overview", useCaseRequest(monthOverviewHandler)).Methods("GET")
  
  r.HandleFunc("/projects", useCaseRequest(projectslistHandler)).Methods("GET")
  r.HandleFunc("/projects/upsert", useCaseRequest(projectUpsertHandler)).Methods("GET")
  r.HandleFunc("/projects/upsert", useCaseRequest(projectUpsertPostHandler)).Methods("POST")
  r.HandleFunc("/projects/upsert/{id}", useCaseRequest(projectUpsertHandler)).Methods("GET")
  r.HandleFunc("/projects/upsert/{id}", useCaseRequest(projectUpsertPostHandler)).Methods("POST")
  r.HandleFunc("/projects/delete/{id}", useCaseRequest(projectDeleteHandler)).Methods("GET")
  r.HandleFunc("/project/{id}", useCaseRequest(projectHandler)).Methods("GET")
  r.HandleFunc("/project/{id}", useCaseRequest(projectPostHandler)).Methods("POST")
  r.HandleFunc("/project/{project_id}/item/{item_id}", useCaseRequest(projectItemHandler)).Methods("GET")
  r.HandleFunc("/project/{project_id}/item/{item_id}", useCaseRequest(projectItemPostHandler)).Methods("POST")
  r.HandleFunc("/project/{project_id}/item/{item_id}/toggle", useCaseRequest(togglePostHandler)).Methods("POST")
  r.HandleFunc("/project/{project_id}/item/{item_id}/delete", useCaseRequest(dayItemDeleteHandler)).Methods("GET")
  
  r.HandleFunc("/about", useCaseRequest(aboutHander)).Methods("GET")
  r.HandleFunc("/app", handlerBundleApp).Methods("GET")
  r.HandleFunc("/logout", logout).Methods("GET")
  r.HandleFunc("/emberapp", authenticate(handlerBundle)).Methods("GET")
  // Everything else fails.
  //r.HandleFunc("/{path:.*}", pageNotFound)
  http.Handle("/", r)  
}