package fuzzyhipster

import (
  "fmt"
  "net/http"
  "src/usecases"
  "time"
  "log"
  "errors"
  "encoding/json"
  "github.com/gorilla/mux"
  "github.com/dgrijalva/jwt-go"
)

type TokenResponseJSON struct {
  Username string  `json:"username"`
  Token string     `json:"token"`
}

func init() {
  r := mux.NewRouter()
  // authentication
    
  r.HandleFunc("/api/1/gettoken", authenticated(handlerGetToken)).Methods("GET")
  r.HandleFunc("/api/1/checktoken", handlerCheckToken).Methods("GET")
  r.HandleFunc("/api/1/token", TokenHandler).Methods("GET")
  r.HandleFunc("/api/1/tokenread", TokenReadHander).Methods("GET")
  
  // projects
  r.HandleFunc("/api/1/projects", useCaseMiddleware(ProjectsHandler)).Methods("GET")
  r.HandleFunc("/api/1/projects/{id}", useCaseMiddleware(ProjectHandler)).Methods("GET")
  r.HandleFunc("/api/1/projects", useCaseMiddleware(CreateProjectHandler)).Methods("POST")
  r.HandleFunc("/api/1/projects/{id}", useCaseMiddleware(UpdateProjectHandler)).Methods("PUT")
  r.HandleFunc("/api/1/projects/{id}", DeleteListHandler).Methods("DELETE")
  
  // project lines
  r.HandleFunc("/api/1/projects/{project_id}/lines",  useCaseMiddleware(CreateProjectLineHandler)).Methods("POST")
  r.HandleFunc("/api/1/projects/{project_id}/lines/{id}",  useCaseMiddleware(UpdateProjectLineHandler)).Methods("PUT")
  r.HandleFunc("/api/1/projects/{project_id}/lines/{id}", useCaseMiddleware(DeleteProjectLineHandler)).Methods("DELETE")
  
  // day items
  r.HandleFunc("/api/1/day/{day_id}", useCaseMiddleware(DayHandler)).Methods("GET")
  //r.HandleFunc("/api/1/dayitem/{id}", useCaseMiddleware(ProjectHandler)).Methods("GET")
  r.HandleFunc("/api/1/dayitem", useCaseMiddleware(CreateDayItemHandler)).Methods("POST")
  //r.HandleFunc("/api/1/dayitem/{id}", useCaseMiddleware(UpdateProjectHandler)).Methods("PUT")
  //r.HandleFunc("/api/1/dayitem/{id}", DeleteListHandler).Methods("DELETE")
  
  // lists
  r.HandleFunc("/api/1/lists", ListsHandler).Methods("GET")
  r.HandleFunc("/api/1/lists/{id}", ListHandler).Methods("GET")
  r.HandleFunc("/api/1/lists", CreateListHandler).Methods("POST")
  r.HandleFunc("/api/1/lists/{id}", UpdateListHandler).Methods("PUT")
  r.HandleFunc("/api/1/lists/{id}", DeleteListHandler).Methods("DELETE")
  
  // items
  r.HandleFunc("/api/1/items", ItemsHandler).Methods("GET")
  r.HandleFunc("/api/1/items/{id}", ItemHandler).Methods("GET")
  
  
  r.HandleFunc("/app", handlerBundleApp).Methods("GET")
  r.HandleFunc("/logout", logout).Methods("GET")
  r.HandleFunc("/", authenticate(handlerBundle)).Methods("GET")
  // Everything else fails.
  //r.HandleFunc("/{path:.*}", pageNotFound)
  http.Handle("/", r)
  
  InitList()
}

func handlerBundle(w http.ResponseWriter, r *http.Request) {
  htmlPage := bundle()
  fmt.Fprint(w, htmlPage)
}

func handlerBundleApp(w http.ResponseWriter, r *http.Request) {
  htmlPage := bundleJavascript()
  fmt.Fprint(w, htmlPage)
}

// logout
func logout(w http.ResponseWriter, r *http.Request) {
  // remove any token cookie
  http.SetCookie(w, &http.Cookie{
		Name:       "token",
		Value:      "",
		Path:       "/",
    RawExpires: "0",
	})
  
  // log out
  useCases := usecases.NewInteractors(r, "")
  if useCases.User.IsLoggedIn() {
    url, err := useCases.User.LogoutUrl()
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    w.Header().Set("Location", url)
    w.WriteHeader(http.StatusFound)
    return
  } else {
    w.Header().Set("Location", "/")
    w.WriteHeader(http.StatusFound)
    return
  }
}

// make sure the request is authenticaed, if not redirect to login
func authenticate(handler http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Println("Authenticate")
        useCases := usecases.NewInteractors(r, "")
        if !useCases.User.IsLoggedIn() {
  	      url, err := useCases.User.LoginUrl()
          if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
          }
          w.Header().Set("Location", url)
          w.WriteHeader(http.StatusFound)
          return
        } else {
          // save the authenticated token as a cookie
          tokenString, _ := createToken(w, r)
          http.SetCookie(w, &http.Cookie{
		        Name:       "token",
		        Value:      tokenString,
		        Path:       "/",
		        RawExpires: "0",
	        })
          
          // run the standard handler
          handler(w, r)
        }
    }
}

// check that this request is session authenticated
func authenticated(handler http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    // check the request session is authenticated
    useCases := usecases.NewInteractors(r, "")
    if !useCases.User.IsLoggedIn() {
      w.WriteHeader(http.StatusUnauthorized)
      return
    }
    // run the standard handler
    handler(w, r)
  }
}

func handlerGetToken(w http.ResponseWriter, r *http.Request) {
  tokenString, err := createToken(w, r)
  if err != nil {
    log.Println("handlerGetToken, not logged in")
    w.WriteHeader(http.StatusUnauthorized)
    return
  }
  useCases := usecases.NewInteractors(r, "")
  
  j, err := json.Marshal(TokenResponseJSON{Username: useCases.User.Current().Email, Token: tokenString})
  if err != nil {
    panic(err)
  }
  w.Write(j)
}

func handlerCheckToken(w http.ResponseWriter, r *http.Request) {
  
  authority := r.Header.Get("Authorization-Token")
  
  // read the authorisation token
  if authority == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
  
  log.Println(authority)
  
  username := decodeTokenUsername(authority)
  
  log.Println(username)
  
  j, err := json.Marshal(TokenResponseJSON{Token: username})
  if err != nil {
    panic(err)
  }
  w.Write(j)
}

func decodeTokenUsername(tokenString string) string {
    
  const sample = "something"
  key := []byte(sample)
  
  token, err := jwt.Parse(tokenString, func(token *jwt.Token) ([]byte, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return key, nil
	})
  
  // branch out into the possible error from signing
	switch err.(type) {
 
	case nil: // no error
 
		if !token.Valid { // but may still be invalid
			//w.WriteHeader(http.StatusUnauthorized)
			//fmt.Fprintln(w, "WHAT? Invalid Token? F*** off!")
			return ""
		}
 
	case *jwt.ValidationError: // something was wrong during the validation
		vErr := err.(*jwt.ValidationError)
 
		switch vErr.Errors {
		case jwt.ValidationErrorExpired:
			//w.WriteHeader(http.StatusUnauthorized)
			//fmt.Fprintln(w, "Token Expired, get a new one.")
			return ""
 
		default:
			//w.WriteHeader(http.StatusInternalServerError)
			//fmt.Fprintln(w, "Error while Parsing Token!")
			//log.Printf("ValidationError error: %+v\n", vErr.Errors)
			return ""
		}
 
	default: // something else went wrong
		//w.WriteHeader(http.StatusInternalServerError)
		//fmt.Fprintln(w, "Error while Parsing Token!")
		//log.Printf("Token parse error: %v\n", err)
		return ""
	}
  
  username, _ := token.Claims["username"].(string)
  return username
  //fmt.Fprint(w, foo)
}

func decodeTokenNamespace(tokenString string) string {
    
  const sample = "something"
  key := []byte(sample)
  
  token, err := jwt.Parse(tokenString, func(token *jwt.Token) ([]byte, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return key, nil
	})
  
  // branch out into the possible error from signing
	switch err.(type) {
 
	case nil: // no error
 
		if !token.Valid { // but may still be invalid
			//w.WriteHeader(http.StatusUnauthorized)
			//fmt.Fprintln(w, "WHAT? Invalid Token? F*** off!")
			return ""
		}
 
	case *jwt.ValidationError: // something was wrong during the validation
		vErr := err.(*jwt.ValidationError)
 
		switch vErr.Errors {
		case jwt.ValidationErrorExpired:
			//w.WriteHeader(http.StatusUnauthorized)
			//fmt.Fprintln(w, "Token Expired, get a new one.")
			return ""
 
		default:
			//w.WriteHeader(http.StatusInternalServerError)
			//fmt.Fprintln(w, "Error while Parsing Token!")
			//log.Printf("ValidationError error: %+v\n", vErr.Errors)
			return ""
		}
 
	default: // something else went wrong
		//w.WriteHeader(http.StatusInternalServerError)
		//fmt.Fprintln(w, "Error while Parsing Token!")
		//log.Printf("Token parse error: %v\n", err)
		return ""
	}
  
  username, _ := token.Claims["namespace"].(string)
  return username
  //fmt.Fprint(w, foo)
}


func createToken(w http.ResponseWriter, r *http.Request) (string, error) {
  
  useCases := usecases.NewInteractors(r, "")
  if !useCases.User.IsLoggedIn() {
  	err := errors.New("not logged in")
    return "", err
  }
  
  // get the user name and details
  user := useCases.User.Current()
  
  const sample = "something"
  
  token := jwt.New(jwt.GetSigningMethod("HS256"))
  token.Claims["username"] = user.Email
  token.Claims["namespace"] =  user.Id
  token.Claims["exp"] = time.Now().Add(time.Minute * 10).Unix()
  
  key := []byte(sample)
  
  tokenString, _ := token.SignedString(key)
  
  return tokenString, nil
}

func TokenHandler(w http.ResponseWriter, r *http.Request) {
  
  useCases := usecases.NewInteractors(r, "")
  if !useCases.User.IsLoggedIn() {
  	_, err := useCases.User.LoginUrl()
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    w.WriteHeader(http.StatusUnauthorized)
    return
  }
  
  // get the user name and details
  user := useCases.User.Current()
  
  const sample = "something"
  
  token := jwt.New(jwt.GetSigningMethod("HS256"))
  token.Claims["username"] = user.Email
  token.Claims["exp"] = time.Now().Add(time.Minute * 10).Unix()
  
  key := []byte(sample)
  
  tokenString, _ := token.SignedString(key)
  
  http.SetCookie(w, &http.Cookie{
		Name:       "token",
		Value:      tokenString,
		Path:       "/",
		RawExpires: "0",
	})
  
  fmt.Fprint(w, tokenString)
}

func getToken(w http.ResponseWriter, r *http.Request) (string, error) {
  
  tokenCookie, err := r.Cookie("token")
	switch {
	case err == http.ErrNoCookie:
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "No Token, no fun!")
		return "", err
	case err != nil:
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error while Parsing cookie!")
		log.Printf("Cookie parse error: %v\n", err)
		return "", err
	}
 
	// just for the lulz, check if it is empty.. should fail on Parse anyway..
	if tokenCookie.Value == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "No Token, no fun!")
		return "", err
	}
  
  return tokenCookie.Value, nil
  
}

func TokenReadHander(w http.ResponseWriter, r *http.Request) {
  
  tokenCookie, err := getToken(w, r)
  if (err != nil) {
    return
  }
  
  const sample = "something"
  key := []byte(sample)
  
  token, err := jwt.Parse(tokenCookie, func(token *jwt.Token) ([]byte, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return key, nil
	})
  
  // branch out into the possible error from signing
	switch err.(type) {
 
	case nil: // no error
 
		if !token.Valid { // but may still be invalid
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "WHAT? Invalid Token? F*** off!")
			return
		}
 
	case *jwt.ValidationError: // something was wrong during the validation
		vErr := err.(*jwt.ValidationError)
 
		switch vErr.Errors {
		case jwt.ValidationErrorExpired:
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "Token Expired, get a new one.")
			return
 
		default:
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Error while Parsing Token!")
			log.Printf("ValidationError error: %+v\n", vErr.Errors)
			return
		}
 
	default: // something else went wrong
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error while Parsing Token!")
		log.Printf("Token parse error: %v\n", err)
		return
	}
  
  foo, _ := token.Claims["username"]
  
  fmt.Fprint(w, foo)
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


