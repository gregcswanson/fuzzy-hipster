package interfaces

import (
	"src/domain"
	"net/http"
	"appengine"
	"appengine/datastore"
  "log"
)

type ProjectRepository BaseRepository

func NewProjectRepository(request *http.Request) *ProjectRepository {
	projectRepository := new(ProjectRepository)
	projectRepository.request = request
	return projectRepository
}

func (repository *ProjectRepository) Get(ID string)(domain.Project, error) {
  var project domain.Project

  // create the namespace context
  globalContext := appengine.NewContext(repository.request)
  c, _ := appengine.Namespace(globalContext, repository.namespace)
  
  // get the key
  key, err := datastore.DecodeKey(ID)
	if err != nil {
	  return project, err
	}
  // retrieve the project
  	err = datastore.Get(c, key, &project);
	project.ID = ID
  
  return project, err
}

func (repository *ProjectRepository) Store(item domain.Project) (domain.Project, error) {
	// upsert operation
	globalContext := appengine.NewContext(repository.request)
  c, _ := appengine.Namespace(globalContext, repository.namespace)
  if item.ID != "" {
		// update
		key , err := datastore.DecodeKey(item.ID)
		if err != nil {
			return item, err
		}
		_, err = datastore.Put(c, key, &item)
    	if err != nil {
			return item, err
		}
	} else {
		// new
		key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Projects", nil), &item)
    	if err != nil {
        	return item, err
    	} else {
    		item.ID = key.Encode()
    	}
	}
	return item, nil
}

func (repository *ProjectRepository) Delete(item domain.Project) error {
  globalContext := appengine.NewContext(repository.request)
  c, _ := appengine.Namespace(globalContext, repository.namespace)
  
  key , err := datastore.DecodeKey(item.ID)
	if err != nil {
		return err
	}
	err = datastore.Delete(c, key)
  return err
}

func (repository *ProjectRepository) Find(bookID string, active bool) ([]domain.Project, error) {
	var projects []domain.Project
	
  log.Println("ProjectRepository.Find 1")
  
	globalContext := appengine.NewContext(repository.request)
  log.Println("ProjectRepository.Find 2")
  c, errNamespace := appengine.Namespace(globalContext, repository.namespace)
  if errNamespace != nil {
    log.Println(errNamespace)
    return projects, errNamespace
  }
  
  log.Println("ProjectRepository.Find 3")
  q := datastore.NewQuery("Projects").Filter("Active =", active)
	
  
  log.Println("ProjectRepository.Find 4")
  
	keys, err := q.GetAll(c, &projects)
  
  log.Println("ProjectRepository.Find 5")
  
  if err != nil {    
    log.Println("ProjectRepository.Find 6")
    return projects, err
  } else {    
    log.Println("ProjectRepository.Find 7")
    // loop through and add the keys as ID
    for i := 0; i < len(keys); i++ {
      projects[i].ID = keys[i].Encode()
    }
  }
      
    log.Println("ProjectRepository.Find 9")
  return projects, nil
}
