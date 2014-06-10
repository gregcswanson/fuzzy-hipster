package domain

type ListRepository interface {
	Store(item List) (List, error)
  Delete(item List) error
	FindForUser(userid string) ([]List, error)
}

type List struct {
	ID string `datastore:"-"`
	UserID string
	Title string
  Description string
}