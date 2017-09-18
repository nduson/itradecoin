package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"msservices/gateway/public"

	"github.com/gorilla/websocket"
)

type Response struct {
	Result []AskBid `json:"result"`
}

type AskBid struct {
	Market  string  `json:"market"`
	Ask     float64 `json:"ask"`
	Bid     float64 `json:"bid"`
	High    float64 `json:"high"`
	Low     float64 `json:"low"`
	Volume  float64 `json:"volume"`
}
type MainAskBid struct {
	Values []AskBid
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var index []byte

func main() {
	indexFile, err := os.Open("index.html")
	if err != nil {
		fmt.Println(err)
	}
	index, err = ioutil.ReadAll(indexFile)
	if err != nil {
		fmt.Println(err)
	}
	http.HandleFunc("/websocket", ws)
	http.HandleFunc("/", inde)
	http.ListenAndServe(":3000", nil)
}

func inde(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string(index))
}

func ws(w http.ResponseWriter, r *http.Request) {
	//var res []byte

	askBidConstruct := make([]AskBid, 0)
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Client subscribed")

	for {
		time.Sleep(2 * time.Second)

		body, err := public.GetTicker("http://localhost:5052/bittrex_ticker")
		//fmt.Fprint(w, string(body))
		if err != nil {
			//return err
			fmt.Println("Error Getting Response From Bittrex. Going to Our DB For The Rquest Data:", err)
			//res := GetAskBidDB(pair, eID)
			//fmt.Fprint(w, string(res))
			return
		}

		var m interface{}
		err = json.Unmarshal(body, &m)
		if err != nil {
			//panic(err)
			fmt.Println("The error on itradecoin ask and bid:", err)
		}
		t := m.(map[string]interface{})
		for key, val := range t {
			//fmt.Println("Got Key1 As:", key, "||", "Got Values1 As:", val)
			if key == "success" && val == false {
				//panic(err)
				fmt.Println("Got Sucess As False:", val)
			}
			if key == "result" {
				for _, val2 := range val.([]interface{}) {
					//fmt.Println("Got Key2 As:", key2, "||", "Got Values2 As:", val2)
					MarketName := val2.(map[string]interface{})["MarketName"]
					ask := val2.(map[string]interface{})["Ask"]
					bid := val2.(map[string]interface{})["Bid"]
					//last := val2.(map[string]interface{})["Last"]
					high24hr := val2.(map[string]interface{})["High"]
					low24hr := val2.(map[string]interface{})["Low"]
					vol := val2.(map[string]interface{})["Volume"]
					//baseVol := val2.(map[string]interface{})["BaseVolume"]
					//exchangeID := 2

					result := AskBid {
						Market:  MarketName.(string),
						Ask:     ask.(float64),
						Bid:     bid.(float64),
						High:    high24hr.(float64),
						Low:     low24hr.(float64),
						Volume:  vol.(float64),
					}
					askBidConstruct = append(askBidConstruct, result)
					//res, _ = json.Marshal(result)
					//fmt.Println(string(res))
					//fmt.Fprint(w, string(res))
				}
			}
		}
		data := Response {
			Result: askBidConstruct,
		}
		response, _ := json.Marshal(data)
		err = conn.WriteMessage(websocket.TextMessage, response)
		if err != nil {
			fmt.Println(err)
			conn.Close()
			break
		}
		fmt.Println("Client unsubscribed")
	}
}

/* func ws(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Client subscribed")

	myPerson := Person{
		Name: "Bill",
		Age:  0,
	}

	for {
		time.Sleep(2 * time.Second)
		if myPerson.Age < 40 {
			myJson, err := json.Marshal(myPerson)
			if err != nil {
				fmt.Println(err)
				return
			}
			err = conn.WriteMessage(websocket.TextMessage, myJson)
			if err != nil {
				fmt.Println(err)
				break
			}
			myPerson.Age += 2
		} else {
			conn.Close()
			break
		}
	}
	fmt.Println("Client unsubscribed")
} */
