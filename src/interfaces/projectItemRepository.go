package interfaces

import (
	"src/domain"
	"net/http"
	"appengine"
	"appengine/datastore"
  "time"
)

type ProjectItemRepository BaseRepository

func NewProjectItemRepository(request *http.Request) *ProjectItemRepository {
	projectItemRepository := new(ProjectItemRepository)
	projectItemRepository.request = request
	return projectItemRepository
}

func (repository *ProjectItemRepository) Store(item domain.ProjectItem) (domain.ProjectItem, error) {
	// upsert operation
	globalContext := appengine.NewContext(repository.request)
  c, _ := appengine.Namespace(globalContext, repository.namespace)
  
    // Add default values
  if item.Sort == 0 {
    item.Sort = time.Now().Unix()
  }
  
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
		key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "ProjectItems", nil), &item)
    	if err != nil {
        	return item, err
    	} else {
    		item.ID = key.Encode()
    	}
	}
	return item, nil
}

func (repository *ProjectItemRepository) Delete(item domain.ProjectItem) error {
  globalContext := appengine.NewContext(repository.request)
  c, _ := appengine.Namespace(globalContext, repository.namespace)
  
  key , err := datastore.DecodeKey(item.ID)
	if err != nil {
		return err
	}
	err = datastore.Delete(c, key)
  return err
}

func (repository *ProjectItemRepository) Find(projectID string) ([]domain.ProjectItem, error) {
	var projectItems []domain.ProjectItem
	
	globalContext := appengine.NewContext(repository.request)
  c, errNamespace := appengine.Namespace(globalContext, repository.namespace)
  if errNamespace != nil {
    return projectItems, errNamespace
  }
  
  q := datastore.NewQuery("ProjectItems").Filter("ProjectID =", projectID).Order("Sort")
	keys, err := q.GetAll(c, &projectItems)
  if err != nil {    
    return projectItems, err
  } else {    
    // loop through and add the keys as ID
    for i := 0; i < len(keys); i++ {
      projectItems[i].ID = keys[i].Encode()
    }
  }
  return projectItems, nil
}

func (repository *ProjectItemRepository) Get(id string)(domain.ProjectItem, error) {
  var projectItem domain.ProjectItem

  // create the namespace context
  globalContext := appengine.NewContext(repository.request)
  c, _ := appengine.Namespace(globalContext, repository.namespace)
  
  // get the key
  key, err := datastore.DecodeKey(id)
	if err != nil {
	  return projectItem, err
	}
  // retrieve the project
  	err = datastore.Get(c, key, &projectItem);
	
  return projectItem, err
}