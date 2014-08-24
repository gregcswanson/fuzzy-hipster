package interfaces

import (
  "net/http"
)

type BaseRepository struct {
  request    *http.Request
  namespace  string
}

type DomainContext struct {
  Lists           ListRepository
  User            UserRepository
  Projects        ProjectRepository
  ProjectItems    ProjectItemRepository
  DayItems        DayItemRepository
}

func NewDomainContext(request *http.Request, namespace string) *DomainContext {
  domainContext := new(DomainContext)
	
  // setup the user repositorytory)
  domainContext.User = UserRepository{}
	domainContext.User.request = request
	domainContext.User.namespace = namespace
  
  // setup the list repository
  domainContext.Lists = ListRepository{}
	domainContext.Lists.request = request
	domainContext.Lists.namespace = namespace
  
  // setup the projects repository
  domainContext.Projects = ProjectRepository{}
	domainContext.Projects.request = request
	domainContext.Projects.namespace = namespace
  
  // setup the projects repository
  domainContext.ProjectItems = ProjectItemRepository{}
	domainContext.ProjectItems.request = request
	domainContext.ProjectItems.namespace = namespace
  
  // setup the day items repository
  domainContext.DayItems = DayItemRepository{}
	domainContext.DayItems.request = request
	domainContext.DayItems.namespace = namespace
  
	return domainContext
}