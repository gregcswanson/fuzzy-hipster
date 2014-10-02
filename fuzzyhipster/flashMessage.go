package fuzzyhipster

import (
  "net/http"
	"appengine"
  "appengine/memcache"
  "strings"
  "log"
  //"bytes"
)

// support flash messaged between pages using memcache

func flashMessage(r *http.Request, userId string , message string ) {
  c := appengine.NewContext(r)
  // Create an Item
  keys := []string{"flashmessage", userId}
  item := &memcache.Item{
    Key:   strings.Join(keys,"-"),
    Value: []byte(message),
  }
  // Add the item to the memcache, if the key does not already exist
  if err := memcache.Add(c, item); err == memcache.ErrNotStored {
    log.Println("item with key %q already exists", item.Key)
  } else if err != nil {
    log.Println("error adding item: %v", err)
  }
}

func getFlashMessage(r *http.Request, userId string) (string) {
  c := appengine.NewContext(r)
  // Get the item from the memcache
  keys := []string{"flashmessage", userId}
  key := strings.Join(keys,"-")
  if item, err := memcache.Get(c, key); err == memcache.ErrCacheMiss {
      return ""
  } else if err != nil {
      return ""
  } else {
    memcache.Delete(c, key)
    return string(item.Value)
  }
  return ""
}

func flashError(r *http.Request, userId string , message string ) {
  c := appengine.NewContext(r)
  // Create an Item
  keys := []string{"flashmessageerror", userId}
  item := &memcache.Item{
    Key:   strings.Join(keys,"-"),
    Value: []byte(message),
  }
  // Add the item to the memcache, if the key does not already exist
  if err := memcache.Add(c, item); err == memcache.ErrNotStored {
    log.Println("item with key %q already exists", item.Key)
  } else if err != nil {
    log.Println("error adding item: %v", err)
  }
}

func getFlashError(r *http.Request, userId string) (string) {
  c := appengine.NewContext(r)
  // Get the item from the memcache
  keys := []string{"flashmessageerror", userId}
  key := strings.Join(keys,"-")
  if item, err := memcache.Get(c, key); err == memcache.ErrCacheMiss {
      return ""
  } else if err != nil {
      return ""
  } else {
    memcache.Delete(c, key)
    return string(item.Value)
  }
  return ""
}

func flashInfo(r *http.Request, userId string , message string ) {
  c := appengine.NewContext(r)
  // Create an Item
  keys := []string{"flashmessageinfo", userId}
  item := &memcache.Item{
    Key:   strings.Join(keys,"-"),
    Value: []byte(message),
  }
  // Add the item to the memcache, if the key does not already exist
  if err := memcache.Add(c, item); err == memcache.ErrNotStored {
    log.Println("item with key %q already exists", item.Key)
  } else if err != nil {
    log.Println("error adding item: %v", err)
  }
}

func getFlashInfo(r *http.Request, userId string) (string) {
  c := appengine.NewContext(r)
  // Get the item from the memcache
  keys := []string{"flashmessageinfo", userId}
  key := strings.Join(keys,"-")
  if item, err := memcache.Get(c, key); err == memcache.ErrCacheMiss {
      return ""
  } else if err != nil {
      return ""
  } else {
    memcache.Delete(c, key)
    return string(item.Value)
  }
  return ""
}

func flashWarning(r *http.Request, userId string , message string ) {
  c := appengine.NewContext(r)
  // Create an Item
  keys := []string{"flashmessagewarning", userId}
  item := &memcache.Item{
    Key:   strings.Join(keys,"-"),
    Value: []byte(message),
  }
  // Add the item to the memcache, if the key does not already exist
  if err := memcache.Add(c, item); err == memcache.ErrNotStored {
    log.Println("item with key %q already exists", item.Key)
  } else if err != nil {
    log.Println("error adding item: %v", err)
  }
}

func getFlashWarning(r *http.Request, userId string) (string) {
  c := appengine.NewContext(r)
  // Get the item from the memcache
  keys := []string{"flashmessagewarning", userId}
  key := strings.Join(keys,"-")
  if item, err := memcache.Get(c, key); err == memcache.ErrCacheMiss {
      return ""
  } else if err != nil {
      return ""
  } else {
    memcache.Delete(c, key)
    return string(item.Value)
  }
  return ""
}