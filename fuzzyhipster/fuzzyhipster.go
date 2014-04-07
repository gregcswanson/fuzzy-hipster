package fuzzyhipster

import (
    //"fmt"
    "net/http"
    //"io/ioutil"
    "github.com/gorilla/mux"
)



func init() {
    r := mux.NewRouter()
    r.HandleFunc("/api/checklists", CheckListsHandler).Methods("GET")
    //r.HandleFunc("/api/kittens/{id}", KittenHandler).Methods("GET")
    //r.HandleFunc("/api/kittens", KittensHandler).Methods("POST")
    //r.HandleFunc("/api/kittens", KittensHandler).Methods("PUT")
    //r.HandleFunc("/api/kittens", KittensHandler).Methods("DELETE")
    r.HandleFunc("/", handler).Methods("GET")
    // Everything else fails.
    //r.HandleFunc("/{path:.*}", pageNotFound)
    http.Handle("/", r)
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