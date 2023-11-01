package controllers

import (
	"html/template"
	"net/http"

	"github.com/raraykinvalery/blowup/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "Is there a free version?",
			Answer:   "Yes! Free trial for 30 days.",
		},
		{
			Question: "What are your support hours?",
			Answer:   "24/7",
		},
		{
			Question: "How do I contact support?",
			Answer: `Email us -
			<a href="mailto:support@blowup.com">
			support@blowup.com </a>`,
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
