package interfaces

import (
	"src/domain"
	"net/http"
	"appengine"
	"appengine/datastore"
  "time"
  "log"
)

type MonthItemRepository BaseRepository

func NewMonthItemRepository(request *http.Request) *MonthItemRepository {
	monthItemRepository := new(MonthItemRepository)
	monthItemRepository.request = request
	return monthItemRepository
}

func (repository *MonthItemRepository) Store(item domain.MonthItem) (domain.MonthItem, error) {
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
		key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "MonthItems", nil), &item)
    	if err != nil {
        	return item, err
    	} else {
    		item.ID = key.Encode()
    	}
	}
	return item, nil
}

func (repository *MonthItemRepository) Find(year int, month int) ([]domain.MonthItem, error) {
  var monthItems []domain.MonthItem
	
	globalContext := appengine.NewContext(repository.request)
  c, errNamespace := appengine.Namespace(globalContext, repository.namespace)
  if errNamespace != nil {
    return monthItems, errNamespace
  }
  
  q := datastore.NewQuery("MonthItems").Filter("Year = ", year).Filter("Month = ", month)
	keys, err := q.GetAll(c, &monthItems)
  if err != nil {    
    log.Println(err.Error())
    return monthItems, err
  } else {    
    // loop through and add the keys as ID
    for i := 0; i < len(keys); i++ {
      monthItems[i].ID = keys[i].Encode()
    }
  }
  return monthItems, nil
}

func (repository *MonthItemRepository) Get(id string)(domain.MonthItem, error) {
  var monthItem domain.MonthItem

  // create the namespace context
  globalContext := appengine.NewContext(repository.request)
  c, _ := appengine.Namespace(globalContext, repository.namespace)
  
  // get the key
  key, err := datastore.DecodeKey(id)
	if err != nil {
	  return monthItem, err
	}
  // retrieve the project
  err = datastore.Get(c, key, &monthItem);
  monthItem.ID = id
	
  return monthItem, err
}

func (repository *MonthItemRepository) Delete(id string) error {
  globalContext := appengine.NewContext(repository.request)
  c, _ := appengine.Namespace(globalContext, repository.namespace)
  
  key , err := datastore.DecodeKey(id)
	if err != nil {
		return err
	}
	err = datastore.Delete(c, key)
  return err
}