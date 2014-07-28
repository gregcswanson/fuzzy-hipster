package usecases

import (
	"src/interfaces"
  "net/http"
	//"errors"
)

type Interactors struct {
  Lists ListInteractor 
  User  UserInteractor
}

type ListInteractor struct {
	Context interfaces.DomainContext
}

type List struct {
  Title string
}

func NewInteractors(request *http.Request, namespace string) *Interactors {
  context := interfaces.NewDomainContext(request, namespace)
  interactors := new(Interactors)
  
  interactors.User = UserInteractor{}
  interactors.User.Context = *context
  
	interactors.Lists = ListInteractor{}
	interactors.Lists.Context = *context
	return interactors
}

func (interactor *ListInteractor) FindForUser() (List, error) {
//	user := interactor.UserRepository.FindCurrent()
//	domainEcho, err := interactor.EchoRepository.FindForUser(user.Id)
//	if err != nil {
//		return Echo{}, err
//	}
//	echo := Echo{domainEcho.ID, domainEcho.Title, nil}
	
//	if echo.ID != "" {
//		lines, _ := interactor.EchoLineRepository.FindByEchoID(echo.ID)
//		echo.Lines = make([]EchoLine, len(lines))
//		for i := 0; i < len(lines); i++ {
//			echo.Lines[i] = EchoLine{lines[i].ID, lines[i].Name}
//		}
//	} else {
//		echo.Lines = make([]EchoLine, 1)
 // 	 	echo.Lines[0] = EchoLine{"1", "Add a new line"}
//	}
	
  l := List{}
  
	return l, nil
}

