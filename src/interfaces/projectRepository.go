package interfaces

import (
	"src/domain"
	"net/http"
	"appengine"
	"appengine/datastore"
)

type ProjectRepository BaseRepository

func NewProjectRepository(request *http.Request) *ProjectRepository {
	projectRepository := new(ProjectRepository)
	projectRepository.request = request
	return projectRepository
}

func (repository *ProjectRepository) Store(item domain.Project) (domain.Project, error) {
	// upsert operation
	c := appengine.NewContext(repository.request)
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
  return nil
}

func (repository *ProjectRepository) Find(bookID string, active bool) ([]domain.Project, error) {
	var projects []domain.Project
	
	globalContext := appengine.NewContext(repository.request)
  c, _ := appengine.Namespace(globalContext, repository.namespace)
  q := datastore.NewQuery("Projects").Filter("Active =", active).Limit(1)
	
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
