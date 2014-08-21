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
	
  
	globalContext := appengine.NewContext(repository.request)
  c, errNamespace := appengine.Namespace(globalContext, repository.namespace)
  if errNamespace != nil {
    log.Println(errNamespace)
    return projects, errNamespace
  }
  
  q := datastore.NewQuery("Projects").Filter("Active =", active).Order("Title")
	
  
	keys, err := q.GetAll(c, &projects)
  
  
  if err != nil {    
    return projects, err
  } else {    
    // loop through and add the keys as ID
    for i := 0; i < len(keys); i++ {
      projects[i].ID = keys[i].Encode()
    }
  }
      
  return projects, nil
}
