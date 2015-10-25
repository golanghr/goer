package main

import (
	"log"
	"text/template"
)

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
