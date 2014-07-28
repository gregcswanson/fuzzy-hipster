package domain

/*
Define the interfaces for the repositories [SPLIT TO DIFFERENT FILE]
Using Verbs Store Find

*/

type UserRepository interface {
	FindCurrent() User
	Store(user User) User
  LoginUrl() (string, error)
  LogoutUrl() (string, error)
}

type User struct {
	Id   string
	Name string
  Email string
  Nickname string
  IsLoggedIn bool
}
