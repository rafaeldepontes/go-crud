package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"rafael.br/simple-crud/internal/handler"
	log "github.com/sirupsen/logrus"
)

const banner = 
`
                Simple CRUD Service                 
 ----------------------------------------------------
   Status  : Starting up
   Version : 1.0.0
   Go      : runtime.GoVersion()
 ----------------------------------------------------
`

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handler.Handler(r)

	fmt.Println(banner)

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Error(err)
	}
}