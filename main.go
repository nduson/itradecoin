package main

import (
	"fmt"
	"msservices/microservices/bittrex"
	"net/http"
	"time"
)

func main() {

	marketData()

	http.ListenAndServe(":5050", nil)
}

func marketData() {
	for {
		go bittrex.BittrexMarketData()

		time.Sleep(5 * time.Second)
	}
}

func init() {

	var name = "iTradeCoin Micro-Services"
	var version = "0.001 DEVEL"
	var developer = "iYochu Nig LTD"

	fmt.Println("App Name: ", name)
	fmt.Println("App Version: ", version)
	fmt.Println("Developer Name: ", developer)
}
