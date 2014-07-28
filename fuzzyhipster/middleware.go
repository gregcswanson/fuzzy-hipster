package fuzzyhipster

import (
  "log"
  "net/http"
  "src/usecases"
)

// Usages include authentication and setup of use cases / domain contexts for routes
// required layers
// 1. Authentication - get the user from the authentication or JWT
// 2. Shared Data Use Cases
// 3. Tenant data Use Cases


/// Usage http.HandleFunc("/hello", basicMiddleware(helloHandler))
func basicMiddleware(handler http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        host := r.Host
      
        //if no active user session then authorize user
		    //if err != nil || user.Id() == "" {
			  //  http.Redirect(w, r, Config.LoginRedirect, http.StatusSeeOther)
			  //  return
		    //}
      
        log.Println(host)
        handler(w, r)
    }
}

func chainMiddleware(handler http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Println("Chained")
        handler(w, r)
    }
}

type UseCaseHandler func(http.ResponseWriter, *http.Request, *usecases.Interactors) 

// check that the request has a valid token in the header and create the 
// use case interactor
func useCaseMiddleware(handler UseCaseHandler) http.HandlerFunc {
   return func(w http.ResponseWriter, r *http.Request) {
      // check the request header for the token
      authority := r.Header.Get("Authorization-Token")

      // read the authorisation token
      if authority == "" {
        w.WriteHeader(http.StatusUnauthorized)
        return
      } 
      username := decodeTokenUsername(authority)
      // create the use case service
      useCases := usecases.NewInteractors(r, username)
      handler(w, r, useCases)
    }
}

