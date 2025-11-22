package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"rafael.br/simple-crud/internal/repository"
)

type BalanceParams struct {
	Username string `json:"username"`
}

type BalanceResponse struct {
	Code    uint16  `json:"code"`
	Balance float64 `json:"balance"`
}

type Error struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	TimeStamp string `json:"timestamp"`
}

const BrazilianDateFormat = "02/01/2006 15:04:05"

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An unexpected Error Occurred.", http.StatusInternalServerError)
	}
)

func writeError(w http.ResponseWriter, msg string, code int) {
	var timestamp string = time.Now().Format(BrazilianDateFormat)
	resp := Error{
				Code: code,
				Message: msg,
				TimeStamp: timestamp,
			}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

func Init() (*sql.DB, error) {
	db, err := repository.Open()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return db, nil
}