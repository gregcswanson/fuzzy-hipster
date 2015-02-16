package interfaces

import (
	"src/domain"
	"net/http"
	"appengine"
	"appengine/datastore"
  "errors"
)

type MonthSummaryRepository BaseRepository

func NewMonthSummaryRepository(request *http.Request) *MonthSummaryRepository {
	monthSummaryRepository := new(MonthSummaryRepository)
	monthSummaryRepository.request = request
	return monthSummaryRepository
}

func (repository *MonthSummaryRepository) Store(monthSummary domain.MonthSummary) (domain.MonthSummary, error) {
	// upsert operation
	globalContext := appengine.NewContext(repository.request)
  c, _ := appengine.Namespace(globalContext, repository.namespace)
  
  if monthSummary.ID != "" {
    //log.Println("Update Day Record")
    //log.Println(item.ID)
		// update
		key , err := datastore.DecodeKey(monthSummary.ID)
		if err != nil {
			return monthSummary, err
		}
    //log.Println(key)
		_, err = datastore.Put(c, key, &monthSummary)
    	if err != nil {
			return monthSummary, err
		}
	} else {
    //log.Println("Create Day Record")
		// new
		key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "MonthSummarys", nil), &monthSummary)
    	if err != nil {
        	return monthSummary, err
    	} else {
    		monthSummary.ID = key.Encode()
    	}
	}
	return monthSummary, nil
}

func (repository *MonthSummaryRepository) FindMonth(year int, month int) (domain.MonthSummary, error) {
  var monthSummary domain.MonthSummary
  var monthSummaries []domain.MonthSummary
	
	globalContext := appengine.NewContext(repository.request)
  c, errNamespace := appengine.Namespace(globalContext, repository.namespace)
  if errNamespace != nil {
    return monthSummary, errNamespace
  }
  
  q := datastore.NewQuery("MonthSummarys").Filter("year = ", year).Filter("Month = ", month)
	keys, err := q.GetAll(c, &monthSummaries)
  if err != nil {
    return monthSummary, err
  } 
  if len(keys) < 1 {
    return monthSummary, errors.New("Record not found")
  }
  
  // loop through and add the key as ID
  monthSummaries[0].ID = keys[0].Encode()
  monthSummary = monthSummaries[0]
  return monthSummary, nil
}

func (repository *MonthSummaryRepository) Delete(id string) error {
  globalContext := appengine.NewContext(repository.request)
  c, _ := appengine.Namespace(globalContext, repository.namespace)
  
  key , err := datastore.DecodeKey(id)
	if err != nil {
		return err
	}
	err = datastore.Delete(c, key)
  return err
}