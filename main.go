package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/raraykinvalery/blowup/views"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	HTMLTpl, err := views.Parse(filepath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(
			w,
			"There was a error parsing the template",
			http.StatusInternalServerError,
		)
		return
	}
	HTMLTpl.Execute(w, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "templates/home.gohtml")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "templates/contact.gohtml")
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
