package domain

//import {
//  "time"
//}

type ProjectRepository interface {
	Store(project Project) (Project, error)
  Delete(project Project) error
	Find(bookID string, active bool) ([]Project, error)
}

type Project struct {
	ID string `datastore:"-"`
  BookID string
	Title string
  Count int
  Open int
  Active bool
  //Created Time
  //Closed Time
}