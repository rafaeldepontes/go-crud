package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
	"rafael.br/simple-crud/api"
	"rafael.br/simple-crud/internal/tool"
)

func GetBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.BalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tool.DatabaseInterface
	database, err = tool.NewDataBase()
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails *tool.BalanceDetails = (*database).GetUserBalance(params.Username)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	resp := api.BalanceResponse{
		Code: http.StatusOK,
		Balance: tokenDetails.Balance,
	}

	w.Header().Set("Content-Type", "application/json")
	
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}