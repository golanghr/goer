package main

import (
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/golanghr/goer/lib/handlers/gzip"
	"github.com/golanghr/goer/lib/handlers/static"
)

const sitePath = "./"
const httpPort = "9000"

var tpls map[string]*template.Template

func init() {
	tpls = templateList(sitePath + "templates/")
}

func main() {
	http.Handle("/", &registration{})

	http.Handle("/resources/", gzip.Handler(static.Handler(sitePath)))

	s := &http.Server{
		Addr:         ":" + httpPort,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	log.Println("Server started on port " + httpPort + "...")
	log.Print(s.ListenAndServe())
}
