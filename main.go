package main

import (
	"html/template"
	"log"
	"net/http"
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
	regTxtStorage := &RegTxtStorage{Filename: sitePath + "registrations"}

	http.Handle("/", &Registration{
		Storer:   regTxtStorage,
		InfoFile: sitePath + "content/event_info.mk"})

	http.Handle("/resources/", gzip.Handler(static.Handler(sitePath)))

	s := &http.Server{
		Addr:         ":" + httpPort,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	log.Println("Server started on port " + httpPort + "...")
	log.Print(s.ListenAndServe())
}

func templateList(templateDir string) (templeteList map[string]*template.Template) {
	tpls := make(map[string]*template.Template)

	driver := template.Must(template.New("master.html").ParseFiles(templateDir + "master.html"))

	list := []string{
		"registration.html",
	}

	for _, tplName := range list {
		subDriver, err := driver.Clone()
		if err != nil {
			log.Fatal("cloning template: ", err)
		}
		_, err = subDriver.ParseFiles(templateDir + tplName)
		if err != nil {
			log.Fatal("parsing ", tplName, ": ", err)
		}

		// strip .html sufix
		keyName := tplName[:len(tplName)-5]

		tpls[keyName] = subDriver
	}

	return tpls
}
