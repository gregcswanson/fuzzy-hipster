package usecases

import (
	"src/interfaces"
  "src/domain"
  "time"
  "errors"
  "log"
)

type Project struct {
	ID      string
	Title   string
  Lines  []ProjectLine
}

type ProjectLine struct {
	ID string
  Status string // open, closed, note, cancelled, running
	Text string
}

type ProjectInteractor struct {
	Context interfaces.DomainContext
}

func (interactor *ProjectInteractor) FindActive() ([]Project, error) {
	// get the active projects
  log.Println("usecases.FindActive")
  activeProjects, _ := interactor.Context.Projects.Find("", true)
	// Copy to the use case model
  var projects []Project
	projects = make([]Project, len(activeProjects))
	for i, project := range activeProjects {
    projectLines := make([]ProjectLine, 0)
		projects[i] = Project{project.ID, project.Title,  projectLines}
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

func (interactor *ProjectInteractor) FindByID(id string) (Project, error) {
  // get the project
  log.Println("usecases.FindByID")
  log.Println(id)
  domainProject, err := interactor.Context.Projects.Get(id)
  if err != nil {
    return Project{}, err
  }   
  // get the lines for the project
  lines, _ := interactor.Context.ProjectItems.Find(id)
	
  // Copy to the use case model
  var projectLines []ProjectLine
	projectLines = make([]ProjectLine, len(lines))
	for i, projectLine := range lines {
		projectLines[i] = ProjectLine{projectLine.ID, projectLine.Status, projectLine.Text}
	}
  
	// Copy to the use case model
  project := Project{domainProject.ID, domainProject.Title, projectLines}
	return project, nil
}

func (interactor *ProjectInteractor) SaveItem(id string, line ProjectLine) (ProjectLine, error) {
  // validate 
	if line.Text == "" {
		err := errors.New("Text is required")
		return line, err
	}
  
  if line.Status == "" {
    line.Status = "NOTE"
  }
	
  // either save or create
	entity := domain.ProjectItem{}
  if line.ID != "" {
    // get the current entity
    entity, _ = interactor.Context.ProjectItems.Get(line.ID)
  } else {
    // setup the new record
    entity.ProjectID = id
  }
	entity.Status = line.Status
  entity.Text = line.Text
  
	// save
	storedEntity, err := interactor.Context.ProjectItems.Store(entity)
	if err == nil {
		line.ID = storedEntity.ID
	}
	
	return line, err
}


func (interactor *ProjectInteractor) Delete(id string) (error) {
  // get the project
  
  // delete the lines
  
  // delete the header
  return nil
}

// SaveLine
// DeleteLine


