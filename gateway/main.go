package main

import (
	"fmt"
	"io/ioutil"
	"msservices/gateway/public"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)

func main() {

	//marketData()
	http.HandleFunc("/", index)
	http.HandleFunc("/ticker", public.GetAskBid)
	http.HandleFunc("/websocket", public.Websocket)
	http.HandleFunc("/test", test)
	http.HandleFunc("/stat", stat)
	http.ListenAndServe(":5000", nil)
}

func marketData() {
	fmt.Println("Enter Bittrex Market Service.... The Service Starts In The Next 5 Sec")
	for {

		//	time.Sleep(5 * time.Second)

		//go public.BittrexMarketData()

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

func test(w http.ResponseWriter, r *http.Request) {

	indexFile, err := os.Open("../socket/index.html")
	if err != nil {
		fmt.Println(err)
	}
	index, err := ioutil.ReadAll(indexFile)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, string(index))
}

func index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "iTradeCoin Live Gate Way Service Is Running")
}

func stat(w http.ResponseWriter, r *http.Request) {

	var n, smallest, biggest time.Duration
	x := public.Stat

	for _, v := range x {
		if v > n {
			fmt.Println(v, ">", n)
			n = v
			biggest = n
		} else {
			fmt.Println(v, "<", n)
		}
	}

	fmt.Println("The biggest number is ", biggest)
	for _, v := range x {
		if v > n {
			fmt.Println(v, ">", n)
		} else {
			fmt.Println(v, "<", n)
			n = v
			smallest = n
		}
	}
	fmt.Println("The smallest number is ", smallest)

	fmt.Fprint(w, "iTradeCoin Live Gate Way Service Is Running", "\n\n")
	fmt.Fprint(w, "Number Of Request Recevied Within 1Hr: ", public.RequestNo, "\n")
	fmt.Fprint(w, "Minimum Execution Time Within 1Hr: ", smallest, "\n")
	fmt.Fprint(w, "Maximum Execution Time Within 1Hr: ", biggest, "\n")
}

func init() {

	var name = "iTradeCoin Live Gate Way Service"
	var version = "0.001 DEVEL"
	var developer = "iYochu Nig LTD"

	fmt.Println("App Name: ", name)
	fmt.Println("App Version: ", version)
	fmt.Println("Developer Name: ", developer)

	//go microservices.TruncateMarketData()
	go clearStatData()
}
