package interfaces

import (
	"src/domain"
  "net/http"
  "appengine"
  "appengine/user"
)

type UserRepository BaseRepository

func NewUserRepositiory(request *http.Request) *UserRepository {
	userRepository := new(UserRepository)
	userRepository.request = request
	return userRepository
}

func (repository *UserRepository) FindCurrent() domain.User {
  // is the current user logged in
  c := appengine.NewContext(repository.request)
  u := user.Current(c)
  var user domain.User
  user = domain.User{}
  if u == nil {
    user.Id = "0"
    user.Name = ""
    user.Nickname = ""
    user.Email = ""
    user.IsLoggedIn = false
  } else {
    user.Id = u.ID
    user.Name = ""
    user.Nickname = ""
    user.Email = u.Email
    user.IsLoggedIn = true
  }  
  return user
}

func (repository *UserRepository) Store(user domain.User) domain.User {
  var u = domain.User{}
  // TO DO
  return u
}

func (repository *UserRepository) LoginUrl()  (string, error){
  c := appengine.NewContext(repository.request)
  u := user.Current(c)
    if u == nil {
        url, err := user.LoginURL(c, repository.request.URL.String())
        return url, err
    }
  return "/", nil 
}

func (repository *UserRepository) LogoutUrl()  (string, error){
  c := appengine.NewContext(repository.request)
  u := user.Current(c)
  if u != nil {
    url, err := user.LogoutURL(c, repository.request.URL.String())
    return url, err
  }
  return "/", nil 
}
