package main

import (
	"net/http"
	"os"
	"regexp"
	"time"
)

type Storer interface {
	Store()
}

type Registration struct {
	Name    string
	Surname string
	Email   string
	Created time.Time
}

func (reg *Registration) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		reg.displayForm(w, r)
	case "POST":
		reg.processForm(w, r)
	default:
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
	}
}

func (reg *Registration) displayForm(w http.ResponseWriter, r *http.Request) {
	p := make(map[string]interface{})

	tpls["registration"].Execute(w, p)
}

func (reg *Registration) processForm(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	surname := r.FormValue("surname")
	email := r.FormValue("email")

	if len(name) == 0 ||
		len(surname) == 0 ||
		len(email) == 0 ||
		!validateEmail(email) {
		http.Error(w, "Form submition not valid.", http.StatusBadRequest)
	}

	reg.Name = name
	reg.Surname = surname
	reg.Email = email
}

type RegistrationTxtStorage struct {
	Registration
	Filename string
}

func (r *RegistrationTxtStorage) Store() {
	f, err := os.OpenFile(r.Filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString(r.Created.String() + " | " + r.Name + " | " + r.Surname + " | " + r.Email + "\n"); err != nil {
		panic(err)
	}
}

func (reg *RegistrationTxtStorage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		reg.Registration.displayForm(w, r)
	case "POST":
		reg.Registration.processForm(w, r)
		reg.Store()
	default:
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
	}
}

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)
}
