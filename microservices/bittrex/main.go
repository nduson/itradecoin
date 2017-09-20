package main

import (
	"fmt"
	"msservices/microservices/bittrex/public"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/bittrex_ticker", public.BittrexMarketData)
	http.HandleFunc("/pair/price", public.AskBidPair)
	http.HandleFunc("/stat", public.Statz)
	http.ListenAndServe(":5052", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "iTradeCoin Bittrex Market Update Service Is Running On Port 5052")
}

func marketData() {
	fmt.Println("Bittrex Market Data Service Started.... The Service Starts In The Next 5 Sec")
	for {

		time.Sleep(2 * time.Second)

		//go public.BittrexMarketDataService()

		fmt.Println("Bittrex Market Service Run Successfully.... The Service Will Run Again In The Next 5 Sec")

	}
}

func clearStatData() {
	fmt.Println("Enter clearStatData.... The Service Starts In The Next 5 Sec")
	for {

		time.Sleep(3600 * time.Second)
		public.RequestNo = 0
		public.Stat = nil

		fmt.Println("clearStatData Run Successfully.... The Service Will Run Again In The Next 5 Sec")

	}
}

func init() {

	var name = "iTradeCoin Bittrex Market Update Service"
	var version = "0.001 DEVEL"
	var developer = "iYochu Nig LTD"

	fmt.Println("App Name: ", name)
	fmt.Println("App Version: ", version)
	fmt.Println("Developer Name: ", developer)

	// start bittrex market data service
	//marketData()
	//clearStatData()
}
