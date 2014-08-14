package usecases

import (
	"src/interfaces"
  "src/domain"
  "time"
  "errors"
  //"log"
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
  //log.Println("usecases.FindActive 1")
  activeProjects, _ := interactor.Context.Projects.Find("", true)
  //log.Println("usecases.FindActive 2")
	// Copy to the use case model
  var projects []Project
	projects = make([]Project, len(activeProjects))
  //log.Println("usecases.FindActive 3")
	for i, project := range activeProjects {
		projects[i] = Project{project.ID, project.Title}
	}
  //log.Println("usecases.FindActive 4")
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
    entity, _ = interactor.Context.Projects.Get(project.ID)
  } else {
    // setup the new record
    entity.Count = 0
    entity.Open = 0
    entity.Active = true
    entity.Start = time.Now()
    entity.End = time.Now()
    entity.BookID = ""
  }
	entity.Title = project.Title
  
	// save
	storedEntity, err := interactor.Context.Projects.Store(entity)
	if err == nil {
		project.ID = storedEntity.ID
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


