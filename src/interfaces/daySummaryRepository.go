package interfaces

import (
	"src/domain"
	"net/http"
	"appengine"
	"appengine/datastore"
	"log"
)

type DaySummaryRepository BaseRepository

func NewDaySummaryRepository(request *http.Request) *DaySummaryRepository {
	daySummaryRepository := new(DaySummaryRepository)
	daySummaryRepository.request = request
	return daySummaryRepository
}

func (repository *DaySummaryRepository) Store(item domain.DaySummary) (domain.DaySummary, error) {
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
		key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "DaySummaries", nil), &item)
    	if err != nil {
        	return item, err
    	} else {
    		item.ID = key.Encode()
    	}
	}
	return item, nil
}

func (repository *DaySummaryRepository) FindForMonth(year int, month int) ([]domain.DaySummary, error) {
  	var daySummaries []domain.DaySummary
	
	globalContext := appengine.NewContext(repository.request)
	c, errNamespace := appengine.Namespace(globalContext, repository.namespace)
	if errNamespace != nil {
    	return daySummaries, errNamespace
  	}
  
  	q := datastore.NewQuery("DaySummaries").Filter("Year = ", year).Filter("Month = ", month)
	keys, err := q.GetAll(c, &daySummaries)
  	if err != nil {    
    	log.Println(err.Error())
    	return daySummaries, err
  	} else {    
    	// loop through and add the keys as ID
    	for i := 0; i < len(keys); i++ {
      	daySummaries[i].ID = keys[i].Encode()
    	}
  	}
  	return daySummaries, nil
}

func (repository *DaySummaryRepository) FindForDay(day int) (domain.DaySummary, error) {
	var daySummaries []domain.DaySummary
	var daySummary domain.DaySummary
	
	globalContext := appengine.NewContext(repository.request)
	c, errNamespace := appengine.Namespace(globalContext, repository.namespace)
	if errNamespace != nil {
    	return daySummary, errNamespace
  	}
  
  	q := datastore.NewQuery("DaySummaries").Filter("Day = ", day)
	keys, err := q.GetAll(c, &daySummaries)
  	if err != nil {    
    	log.Println(err.Error())
    	return daySummary, err
  	} else {    
    	// loop through and add the keys as ID
    	for i := 0; i < len(keys); i++ {
      		daySummaries[i].ID = keys[i].Encode()
    	}
  	}
  	if len(daySummaries) == 0 {
  		return daySummary, nil
  	}
  	return daySummaries[0], nil
}

func (repository *DaySummaryRepository) Get(id string)(domain.DaySummary, error) {
  	var daySummary domain.DaySummary

  	// create the namespace context
  	globalContext := appengine.NewContext(repository.request)
  	c, _ := appengine.Namespace(globalContext, repository.namespace)
  
  	// get the key
  	key, err := datastore.DecodeKey(id)
	if err != nil {
		return daySummary, err
	}
  	
  	// retrieve the project
  	err = datastore.Get(c, key, &daySummary);
  	daySummary.ID = id
	
  	return daySummary, err
}

func (repository *DaySummaryRepository) Delete(id string) error {
	globalContext := appengine.NewContext(repository.request)
	c, _ := appengine.Namespace(globalContext, repository.namespace)
  
	key , err := datastore.DecodeKey(id)
	if err != nil {
		return err
	}
	err = datastore.Delete(c, key)
	return err
}