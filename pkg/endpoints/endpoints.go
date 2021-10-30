package endpoints

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"notes/pkg/models"
	"notes/pkg/service"
	"notes/pkg/tpl"
	"strconv"

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
	r.HandleFunc("/note/{id:.*}", e.GetNote).Methods(http.MethodGet)
	return r
}

type endpoint struct {
	service *service.Service
}

func (e endpoint) GetNotes(w http.ResponseWriter, r *http.Request) {
	notes, err := e.service.GetNotes(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	data := struct {
		Title string
		Items []models.Note
	}{
		Title: "My notes",
		Items: notes,
	}

	tpl.Render(w, r, data, "notesHtml")
}

func (e endpoint) GetNote(w http.ResponseWriter, r *http.Request) {
	url := models.Note{}
	vars := mux.Vars(r)
	id := vars["id"]

	var err error
	url.Id, err = strconv.ParseUint(id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	note, err := e.service.GetNote(r.Context(), url)

	if err != nil {
		log.Println(err)
		if errors.Is(err, models.ErrNoteNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title string
		Item  models.Note
	}{
		Title: "My note",
		Item:  note,
	}

	tpl.Render(w, r, data, "noteHtml")
}
