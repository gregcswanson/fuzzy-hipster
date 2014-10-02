package fuzzyhipster

import (
  "html/template"
  "net/http"
  "reflect"
  "src/usecases"
)

//Compile templates on start
var templates = template.Must(template.ParseFiles(
	"src/templates/header.html", 
	"src/templates/footer.html", 
	"src/templates/index.html", 
	"src/templates/projects.html",
	"src/templates/project/add.html",
    "src/templates/project/index.html",
	"src/templates/about.html")).Funcs(template.FuncMap(map[string]interface{}{"eq": eq}))
 
//A Page structure
type Page struct {
	Title string
	IsDayView bool
	IsMonthView bool
	IsProjectView bool
	IsAboutView bool
	Model interface{}
	Error string
	Info string
	Warning string
	Success string
}

func buildPage(r *http.Request, u *usecases.Interactors) (Page) {
  // get the flash messages
  success := getFlashMessage(r, u.User.Current().Id)
  info := getFlashInfo(r, u.User.Current().Id)
  warning := getFlashWarning(r, u.User.Current().Id)
  error := getFlashError(r, u.User.Current().Id)
  
  page := Page{Title: "Index", Success: success, Warning: warning, Info: info, Error: error }
  return page
}
 
//Render the named template
func render(w http.ResponseWriter, tmpl string, data interface{}) {
	templates.ExecuteTemplate(w, tmpl, data)
}

func eq(args ...interface{}) bool {
        if len(args) == 0 {
                return false
        }
        x := args[0]
        switch x := x.(type) {
        case string, int, int64, byte, float32, float64:
                for _, y := range args[1:] {
                        if x == y {
                                return true
                        }
                }
                return false
        }

        for _, y := range args[1:] {
                if reflect.DeepEqual(x, y) {
                        return true
                }
        }
        return false
}