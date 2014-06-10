package interfaces

import (
	"src/domain"
	"net/http"
	"appengine"
	"appengine/datastore"
)

type ListRepository BaseRepository

func NewListRepository(request *http.Request) *ListRepository {
	listRepository := new(ListRepository)
	listRepository.request = request
	return listRepository
}

func (repository *ListRepository) Store(item domain.List) (domain.List, error) {
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
		key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Lists", nil), &item)
    	if err != nil {
        	return item, err
    	} else {
    		item.ID = key.Encode()
    	}
	}
	return item, nil
}

func (repository *ListRepository) Delete(item domain.List) error {
  return nil
}

func (repository *ListRepository) FindForUser(userid string) ([]domain.List, error) {
	var lists []domain.List
	
	c := appengine.NewContext(repository.request)
  q := datastore.NewQuery("Lists").Filter("UserID =", userid).Limit(1)
	
	keys, err := q.GetAll(c, &lists)
  if err != nil {
    return lists, err
  } else {
    // loop through and add the keys as ID
    for i := 0; i < len(keys); i++ {
      lists[i].ID = keys[i].Encode()
    }
  }
  return lists, nil
}
