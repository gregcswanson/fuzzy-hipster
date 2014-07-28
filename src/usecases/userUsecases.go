package usecases

import (
	"src/interfaces"
  "src/domain"
)

type UserInteractor struct {
	Context interfaces.DomainContext
}

func(interactor *UserInteractor) Current() domain.User {
  user := interactor.Context.User.FindCurrent()
  return user
}

func (interactor *UserInteractor) IsLoggedIn() bool {
  user := interactor.Context.User.FindCurrent()
  return user.IsLoggedIn
}

func (interactor *UserInteractor) LoginUrl() (string, error) {
	url, err := interactor.Context.User.LoginUrl()
	return url, err
}

func (interactor *UserInteractor) LogoutUrl() (string, error) {
	url, err := interactor.Context.User.LogoutUrl()
	return url, err
}