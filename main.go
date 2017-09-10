package main

import (
	"fmt"
	"msservices/microservices"
	"msservices/microservices/bittrex"
	"msservices/microservices/itradecoin"
	"net/http"
	_ "net/http/pprof"
)

func main() {

	marketData()

	http.HandleFunc("/askbid", itradecoin.GetAskBid)
	http.ListenAndServe(":5050", nil)
}

func marketData() {
	//for {
	go bittrex.BittrexMarketData()

	//time.Sleep(100 * time.Second)
	//}
}

func init() {

	var name = "iTradeCoin Micro-Services"
	var version = "0.001 DEVEL"
	var developer = "iYochu Nig LTD"

	fmt.Println("App Name: ", name)
	fmt.Println("App Version: ", version)
	fmt.Println("Developer Name: ", developer)

	go microservices.TruncateMarketData()
}
