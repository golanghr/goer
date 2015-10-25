package main

import "net/http"

type registration struct {
}

func (reg *registration) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := make(map[string]interface{})

	tpls["registration"].Execute(w, p)
}
