package usecases

import (
	"src/interfaces"
  "src/domain"
  "time"
)

type Project struct {
	ID      string
	Title   string
}

type ProjectLine struct {
	ID string
  Status string // open, closed, note, cancelled, in progress
	Name string
}

type ProjectInteractor struct {
	Context interfaces.DomainContext
}

func (interactor *ProjectInteractor) FindActive() ([]Project, error) {
	// get the active projects
  activeProjects := interactor.Context.Projects.Find("", true)
	// Copy to the use case model
  var projects []Project
	projects = make([]Project, len(activeProjects))
	for i, project := range activeProjects {
		projects[i] = Project{project.Id, project.Title}
	}
	return projects, nil
}

func (interactor *ProjectInteractor) Save(project Project) (Project, error) {
	// validate 
	if project.Title == "" {
		err := errors.New("Title is required")
		return project, err
	}
	
  // either save or create
	entity := domain.Project{}
  if project.ID != "" {
    // get the current entity
    entity, err = interactor.Context.Projects.Get(project.ID)
  } else {
    // setup the new record
    entity.Count = 0
    entity.Open = 0
    entity.Active = true
    entity.Open = time.Now()
    entity.End time.Now()
    entity.BookID = ""
  }
	entity.Title = project.Title
  
	// save
	entity, err := interactor.Context.Projects.Store(entity)
	if err == nil {
		project.ID = entity.ID
	}
	
	return project, err
}

func (interactor *ProjectInteractor) Delete(id string) (error) {
  // get the project
  
  // delete the lines
  
  // delete the header
  return nil
}

// SaveLine
// DeleteLine


