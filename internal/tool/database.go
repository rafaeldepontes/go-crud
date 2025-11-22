package tool

import (
	log "github.com/sirupsen/logrus"
)

type LoginDetails struct {
	AuthToken string
	Username string
}
type BalanceDetails struct {
	Balance float64
	Username string
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserBalance(username string) *BalanceDetails
	SetUpDatabase() error
}

func NewDataBase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &mockDB{} // Need to implement this... using a real MySQL or Postgres or whatever...
	var err error = database.SetUpDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}