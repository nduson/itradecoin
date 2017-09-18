package main

import (
	"fmt"
	"msservices/microservices"
	"msservices/microservices/workers/worker"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {

	workers()

	http.HandleFunc("/", index)
	http.ListenAndServe(":6000", nil)
}

func workers() {
	count := 1
	for {
		fmt.Println("Service Have Run This Number Of times ", count)
		timeInterval := microservices.GetTimerInterval("BuyOrderUpdateWorker")
		time.Sleep(timeInterval * time.Second)

		go worker.BuyOrderUpdateWorker()
		go worker.SellOrderUpdateWorker()
		count = count + 1

	}
}

func index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "iTradeCoin Buy Order Worker Service Is Running On Port 6000")
}

func init() {

	var name = "iTradeCoin Buy Order Update Worker"
	var version = "0.001 DEVEL"
	var developer = "iYochu Nig LTD"

	fmt.Println("App Name: ", name)
	fmt.Println("App Version: ", version)
	fmt.Println("Developer Name: ", developer)

	//go microservices.TruncateMarketData()
}
