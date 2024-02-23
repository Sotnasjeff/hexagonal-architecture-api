package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Sotnasjeff/hexagonal-architecture-api/adapters/web/handler"
	"github.com/Sotnasjeff/hexagonal-architecture-api/app"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type Webserver struct {
	Service app.ProductServiceInterface
}

func NewWebServer() *Webserver {
	return &Webserver{}
}

func (w Webserver) Server() {
	r := mux.NewRouter()
	n := negroni.New(negroni.NewLogger())

	handler.MakeProductHandler(r, n, w.Service)
	http.Handle("/", r)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":9000",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log:", log.Lshortfile),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
