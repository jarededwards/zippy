package main

import (
	_ "embed"
	"log"
	"net/http"
	"text/template"
)

// embed the public templating code
//
//go:embed templates/index.html
var code string

var compiledTemplate = template.Must(template.New("index").Parse(code))

func main() {
	mux := http.NewServeMux()

	// This endpoint will show any content in the "public" directory
	// and will also serve the index.html file as the homepage
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		compiledTemplate.Execute(w, nil)
	})

	mux.HandleFunc("/form/hello", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		compiledTemplate.Execute(w, map[string]interface{}{
			"Name": name,
		})
	})

	// I create a server, no framework, just pure Go
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// I start the server, if fails, I log the error
	log.Fatal(srv.ListenAndServe())
}
