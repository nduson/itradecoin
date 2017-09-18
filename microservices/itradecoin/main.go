package main

import (
	"fmt"
	"msservices/microservices/bittrex/public"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {

	marketData()
	http.HandleFunc("/", index)
	http.ListenAndServe(":5051", nil)
}

func marketData() {
	fmt.Println("Enter Bittrex Market Service.... The Service Starts In The Next 5 Sec")
	for {

		time.Sleep(5 * time.Second)

		go public.BittrexMarketData()

		fmt.Println("Bittrex Market Service Run Successfully.... The Service Will Run Again In The Next 5 Sec")

	}
}

func index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "iTradeCoin Bittrex Market Update Service Is Running")
}

func init() {

	var name = "iTradeCoin Bittrex Market Update Service"
	var version = "0.001 DEVEL"
	var developer = "iYochu Nig LTD"

	fmt.Println("App Name: ", name)
	fmt.Println("App Version: ", version)
	fmt.Println("Developer Name: ", developer)

	//go microservices.TruncateMarketData()
}
