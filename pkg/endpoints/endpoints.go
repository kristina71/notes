package endpoints

import (
	"fmt"
	"log"
	"net/http"
	"notes/pkg/models"
	"notes/pkg/service"
	"text/template"

	_ "embed"

	"github.com/gorilla/mux"
)

func New(service *service.Service) http.Handler {
	r := mux.NewRouter()

	e := endpoint{service: service}

	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	}).Methods(http.MethodGet)
	r.HandleFunc("/notes", e.GetNotes).Methods(http.MethodGet)
	return r
}

type endpoint struct {
	service *service.Service
}

//go:embed static/notes.html
var notesTpl string

func (e endpoint) GetNotes(w http.ResponseWriter, r *http.Request) {
	notes, err := e.service.GetNotes(r.Context())
	if err != nil {
		log.Fatal(err)
		return
	}

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("webpage").Parse(notesTpl)
	check(err)

	data := struct {
		Title string
		Items []models.Note
	}{
		Title: "My page",
		Items: notes,
	}

	err = t.Execute(w, data)
	check(err)
}
