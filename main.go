package main

import (
	"log"
	"net/http"
	"time"

	"github.com/golanghr/goer/lib/handlers/gzip"
	"github.com/golanghr/goer/lib/handlers/static"
)

const sitePath = "./"
const httpPort = "9000"

func main() {
	http.Handle("/resources/", gzip.Handler(static.Handler(sitePath)))

	s := &http.Server{
		Addr:         ":" + httpPort,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	log.Println("Server started on port " + httpPort + "...")
	log.Print(s.ListenAndServe())
}
