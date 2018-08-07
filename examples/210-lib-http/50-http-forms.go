package main

import (
	"html/template"
	"net/http"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func contactForm(response http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("forms.html"))

	// check method
	if request.Method != http.MethodPost {
		tmpl.Execute(response, nil)
		return
	}

	details := ContactDetails{
		Email:   request.FormValue("email"),
		Subject: request.FormValue("subject"),
		Message: request.FormValue("message"),
	}

	_ = details // do something

	tmpl.Execute(response, struct{ Success bool }{true})
}

func main() {
	http.HandleFunc("/", contactForm)
	http.ListenAndServe(":8090", nil)
}
