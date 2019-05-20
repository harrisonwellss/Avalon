package main

import (
	"fmt"
	"github.com/okcoin-okex/open-api-v3-sdk/okex-go-sdk-api"
)

func NewOKExClient() *okex.Client {
	var config okex.Config
	config.Endpoint = "https://www.okex.com/"
	config.ApiKey = "15bf8a93-ae53-4b72-951c-cd4ada830d5b"
	config.SecretKey = "C2D121404CB22A6EED2E98D99D53A888"
	config.Passphrase = "123QWEasdzxc"
	config.TimeoutSecond = 45
	config.IsPrint = true
	config.I18n = okex.ENGLISH

	client := okex.NewClient(config)
	return client
}

func main() {

	exchangeRate, err := NewOKExClient().GetSwapCandlesByInstrument("BTC-USD-SWAP", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(exchangeRate)
}
