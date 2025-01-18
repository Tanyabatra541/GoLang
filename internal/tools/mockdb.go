package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"tanya": {
		AuthToken: "tanya123",
		Username: "tanya",
	},
	"nitish": {
		AuthToken: "nitish123",
		Username: "nitish",
	},
	"rakesh": {
		AuthToken: "rakesh123",
		Username: "rakesh",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"tanya": {
		Coins: 100,
		Username: "tanya",
	},
	"nitish": {
		Coins: 200,
		Username: "nitish",
	},
	"rakesh": {
		Coins: 300,
		Username: "rakesh",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)
	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)
	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}

