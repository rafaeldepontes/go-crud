package api

import (
	"encoding/json"
	"net/http"
)

type CoinBalanceParams struct {
	Username string `json:"username"`
}

type CoinBalanceResponse struct {
	Code    uint16  `json:"code"`
	Balance float64 `json:"balance"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An unexpected Error Occurred.", http.StatusInternalServerError)
	}
)

func writeError(w http.ResponseWriter, msg string, code int) {
	resp := Error{
				Code: 401,
				Message: msg,
			}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}