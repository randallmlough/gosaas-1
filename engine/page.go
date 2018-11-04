package engine

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path"
)

var pageTemplates *template.Template

func init() {
	loadTemplates()
}

func loadTemplates() {
	var tmpl []string

	files, err := ioutil.ReadDir("./templates")
	if err != nil {
		log.Fatal("unable to load templates", err)
	}

	for _, f := range files {
		tmpl = append(tmpl, path.Join("./templates", f.Name()))
	}

	t, err := template.ParseFiles(tmpl...)
	if err != nil {
		log.Fatal("error while parsing templates", err)
	}

	pageTemplates = t
}

func ServePage(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	t := pageTemplates.Lookup(name)
	if err := t.Execute(w, data); err != nil {
		fmt.Println("error while rendering the template ", err)
	}

	logRequest(r, http.StatusOK)
}
