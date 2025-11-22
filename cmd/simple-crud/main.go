package main

import (
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"rafael.br/simple-crud/api"
	"rafael.br/simple-crud/internal/handler"
)

const banner = 
`
                Simple CRUD Service                 
 ----------------------------------------------------
   Status  : Starting up
   Version : 1.0.0
   Go      : runtime.GoVersion()
   Listen : localhost:8080
 ----------------------------------------------------
`

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handler.Handler(r)

	db, err := api.Init()
	if err != nil {
		log.Error(err)
	}
	defer db.Close()

	println(banner)

	err = http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Error(err)
	}
}