package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/raraykinvalery/blowup/controllers"
	"github.com/raraykinvalery/blowup/templates"
	"github.com/raraykinvalery/blowup/views"
)

func main() {
	r := chi.NewRouter()

	tpl := views.Must(views.ParseFS(templates.FS, "layout_page.gohtml", "home.gohtml", "layout_parts.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "layout_page.gohtml", "contact.gohtml", "layout_parts.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "layout_page.gohtml", "faq.gohtml", "layout_parts.gohtml"))
	r.Get("/faq", controllers.FAQ(tpl))

	usersC := controllers.Users{}
	usersC.Templates.New = views.Must(
		views.ParseFS(templates.FS,
			"layout_page.gohtml",
			"signup.gohtml",
			"layout_parts.gohtml"),
	)
	r.Get("/signup", usersC.New)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
