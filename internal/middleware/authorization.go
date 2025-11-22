package middleware

import (
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"
	"rafael.br/simple-crud/api"
	tool "rafael.br/simple-crud/internal/tool"
)

var ErrorUnAuthorized = errors.New("Invalid username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var username = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(ErrorUnAuthorized)
			api.RequestErrorHandler(w, ErrorUnAuthorized)
			return
		}

		var database *tool.DatabaseInterface
		database, err = tool.NewDataBase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tool.LoginDetails = (*database).GetUserLoginDetails(username)
		if loginDetails == nil || (token != loginDetails.AuthToken) {
			log.Error(ErrorUnAuthorized)
			api.RequestErrorHandler(w, ErrorUnAuthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}