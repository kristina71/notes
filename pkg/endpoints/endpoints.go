package endpoints

import (
	"fmt"
	"log"
	"net/http"
	"notes/pkg/models"
	"notes/pkg/service"
	"regexp"
	"strconv"
	"strings"
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
	r.HandleFunc("/note/{id:.*}", e.GetNote).Methods(http.MethodGet)
	return r
}

type endpoint struct {
	service *service.Service
}

//go:embed static/notes.html
var notesTpl string

//go:embed static/footer.html
var footer string

//go:embed static/header.html
var header string

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
		Title  string
		Items  []models.Note
		Header string
		Footer string
	}{
		Title:  "My page",
		Items:  notes,
		Header: header,
		Footer: footer,
	}

	err = t.Execute(w, data)
	check(err)
}

//go:embed static/note.html
var noteTpl string

func (e endpoint) GetNote(w http.ResponseWriter, r *http.Request) {
	url := models.Note{}
	str := strings.Trim(r.URL.Path, "/")

	re, _ := regexp.Compile(`\d+`)
	res := re.FindAllString(str, -1)

	value, err := strconv.Atoi(res[0])
	if err != nil {
		log.Fatal(err)
		return
	}
	url.Id = uint16(value)
	fmt.Println(url.Id)

	note, err := e.service.GetNote(r.Context(), url)
	if err != nil {
		log.Fatal(err)
		return
	}

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("webpage").Parse(noteTpl)
	check(err)

	data := struct {
		Title  string
		Item   models.Note
		Header string
		Footer string
	}{
		Title:  "My page",
		Item:   note,
		Header: header,
		Footer: footer,
	}

	err = t.Execute(w, data)
	check(err)
}
