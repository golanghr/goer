package main

import (
	"log"
	"net/http"
	"os"
	"regexp"
	"time"
)

// RegistrationStorer for separation of domain in storing data of registration
type RegistrationStorer interface {
	Store(reg Registration) (err error)
}

// Registration data structure
type Registration struct {
	Name    string
	Surname string
	Email   string
	Created time.Time
	Storer  RegistrationStorer
}

// ServeHTTP bind to http
// Will display form on GET and expects form process on POST
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
	reg.Created = time.Now()

	err := reg.Storer.Store(*reg)
	if err != nil {
		log.Printf("Error %s while storing registration %+v\n", err, reg)
		http.Error(w, "Registration storage failed.", http.StatusInternalServerError)
	}
}

// RegTxtStorage will save registrations in form of basic TXT file
type RegTxtStorage struct {
	Filename string
}

// Store implements RegistrationStorer interface
func (r *RegTxtStorage) Store(reg Registration) (err error) {
	f, err := os.OpenFile(r.Filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.WriteString(reg.Created.String() + " | " + reg.Name + " | " + reg.Surname + " | " + reg.Email + "\n"); err != nil {
		return err
	}

	return nil
}

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)
}
