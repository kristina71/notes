package tpl

import (
	"embed"
	_ "embed"
	"io/fs"
	"log"
	"net/http"
	"text/template"
)

//go:embed static
var Content embed.FS

var tmp *template.Template

func init() {
	tmp = template.New("")
	root, err := fs.Sub(Content, "static")
	if err != nil {
		panic(err)
	}

	tmp, err = tmp.ParseFS(root, "*.html")

	if err != nil {
		panic(err)
	}
}

func Render(w http.ResponseWriter, r *http.Request, data interface{}, tmpName string) {
	err := tmp.ExecuteTemplate(w, tmpName, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
