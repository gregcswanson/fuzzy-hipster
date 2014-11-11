package interfaces

import (
	"src/domain"
	"net/http"
	"appengine"
	"appengine/datastore"
  "time"
  "log"
  "strconv"
  "fmt"
)

type DayItemRepository BaseRepository

func NewDayItemRepository(request *http.Request) *DayItemRepository {
	dayItemRepository := new(DayItemRepository)
	dayItemRepository.request = request
	return dayItemRepository
}

func (repository *DayItemRepository) Store(item domain.DayItem) (domain.DayItem, error) {
	// upsert operation
	globalContext := appengine.NewContext(repository.request)
  c, _ := appengine.Namespace(globalContext, repository.namespace)
  
    // Add default values
  if item.Sort == 0 {
    item.Sort = time.Now().Unix()
  }
  
  if item.ID != "" {
    //log.Println("Update Day Record")
    //log.Println(item.ID)
		// update
		key , err := datastore.DecodeKey(item.ID)
		if err != nil {
			return item, err
		}
    //log.Println(key)
		_, err = datastore.Put(c, key, &item)
    	if err != nil {
			return item, err
		}
	} else {
    //log.Println("Create Day Record")
		// new
		key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "DayItems", nil), &item)
    	if err != nil {
        	return item, err
    	} else {
    		item.ID = key.Encode()
    	}
	}
	return item, nil
}

func (repository *DayItemRepository) Find(dayAsInt int) ([]domain.DayItem, error) {
	var dayItems []domain.DayItem
	
	globalContext := appengine.NewContext(repository.request)
  c, errNamespace := appengine.Namespace(globalContext, repository.namespace)
  if errNamespace != nil {
    return dayItems, errNamespace
  }
  
  q := datastore.NewQuery("DayItems").Filter("Day =", dayAsInt) //.Order("Sort")
	keys, err := q.GetAll(c, &dayItems)
  if err != nil {    
    return dayItems, err
  } else {    
    // loop through and add the keys as ID
    for i := 0; i < len(keys); i++ {
      dayItems[i].ID = keys[i].Encode()
    }
  }
  return dayItems, nil
}

func (repository *DayItemRepository) FindMonth(year int, month int) ([]domain.DayItem, error) {
  var dayItems []domain.DayItem
	
	globalContext := appengine.NewContext(repository.request)
  c, errNamespace := appengine.Namespace(globalContext, repository.namespace)
  if errNamespace != nil {
    return dayItems, errNamespace
  }
  
  // strconv.Atoi(id)
  startOfMonth, _ := strconv.Atoi(fmt.Sprintf("%d%d00", year, month))
  endOfMonth, _ := strconv.Atoi(fmt.Sprintf("%d%d99", year, month))
  
  q := datastore.NewQuery("DayItems").Filter("Day > ", startOfMonth).Filter("Day < ", endOfMonth)
	keys, err := q.GetAll(c, &dayItems)
  if err != nil {    
    log.Println(err.Error())
    return dayItems, err
  } else {    
    // loop through and add the keys as ID
    for i := 0; i < len(keys); i++ {
      dayItems[i].ID = keys[i].Encode()
    }
  }
  return dayItems, nil
}

func (repository *DayItemRepository) Get(id string)(domain.DayItem, error) {
  var dayItem domain.DayItem

  // create the namespace context
  globalContext := appengine.NewContext(repository.request)
  c, _ := appengine.Namespace(globalContext, repository.namespace)
  
  // get the key
  key, err := datastore.DecodeKey(id)
	if err != nil {
	  return dayItem, err
	}
  // retrieve the project
  err = datastore.Get(c, key, &dayItem);
  dayItem.ID = id
	
  return dayItem, err
}

func (repository *DayItemRepository) Delete(id string) error {
  globalContext := appengine.NewContext(repository.request)
  c, _ := appengine.Namespace(globalContext, repository.namespace)
  
  key , err := datastore.DecodeKey(id)
	if err != nil {
		return err
	}
	err = datastore.Delete(c, key)
  return err
}