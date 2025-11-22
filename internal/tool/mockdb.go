package tool

import (
	"time"
	//"database/sql"
)

type mockDB struct {}

var mockLoginDetails = map[string]LoginDetails {
	"alex": {
		AuthToken: "123ABC",
		Username: "alex",
	},
	"jason": {
		AuthToken: "456DEF",
		Username: "jason",
	},
	"rafael": {
		AuthToken: "789GHI",
		Username: "rafael",
	},
}

var mockBalanceDetails = map[string]BalanceDetails {
	"alex": {
		Balance: 129.66,
		Username: "alex",
	},
	"jason": {
		Balance: 9382.21,
		Username: "jason",
	},
	"rafael": {
		Balance: 3817.98,
		Username: "rafael",
	},
}

func (m mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	loginDetails := LoginDetails{}
	loginDetails, ok := mockLoginDetails[username]

	if !ok {
		return nil
	}
	return &loginDetails
}

func (m mockDB) GetUserBalance(username string) *BalanceDetails {
	time.Sleep(time.Second * 1)

	balanceDetails := BalanceDetails{}
	balanceDetails, ok := mockBalanceDetails[username]

	if !ok {
		return nil
	}
	return &balanceDetails
}

func (m mockDB) SetUpDatabase() error {
	return nil
}