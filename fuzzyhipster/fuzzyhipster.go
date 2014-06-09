package fuzzyhipster

import (
    //"fmt"
    "net/http"
    //"io/ioutil"
    "github.com/gorilla/mux"
)


func init() {
    r := mux.NewRouter()
  // lists
    r.HandleFunc("/api/1/lists", ListsHandler).Methods("GET")
    r.HandleFunc("/api/1/lists/{id}", ListHandler).Methods("GET")
    r.HandleFunc("/api/1/lists", CreateListHandler).Methods("POST")
    r.HandleFunc("/api/1/lists/{id}", UpdateListHandler).Methods("PUT")
    r.HandleFunc("/api/1/lists/{id}", DeleteListHandler).Methods("DELETE")
  // items
    r.HandleFunc("/api/1/items", ItemsHandler).Methods("GET")
    r.HandleFunc("/api/1/items/{id}", ItemHandler).Methods("GET")
  
  
    r.HandleFunc("/", handler).Methods("GET")
    // Everything else fails.
    //r.HandleFunc("/{path:.*}", pageNotFound)
    http.Handle("/", r)
  
    InitList()
}

func handler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./public/html/index.html")
}

func pageNotFound(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./public/html/404.html")
}

func authenticateRequest(w http.ResponseWriter, r *http.Request) (string, error) {
  return "steve", nil
}

func bundle() {
  /*

  buf := bytes.NewBuffer(nil)
  for _, filename := range filenames {
    f, _ := os.Open(filename) // Error handling elided for brevity.
    io.Copy(buf, f)           // Error handling elided for brevity.
    f.Close()
  }
  s := string(buf.Bytes())

  */
}

