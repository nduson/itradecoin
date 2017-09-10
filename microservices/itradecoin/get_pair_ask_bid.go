package itradecoin

import (
	"encoding/json"
	"fmt"
	"msservices/microservices/bittrex"
	"net/http"
)

// AskBid is a struct use to return ask and bid of request pair.
type AskBid struct {
	Success string  `json:"success"`
	Message string  `json:"message"`
	Market  string  `json:"market"`
	Ask     float64 `json:"ask"`
	Bid     float64 `json:"bid"`
}

// GetAskBid is the function that will return the ask and bid or error.
func GetAskBid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pair := r.FormValue("pair")
	eID := r.FormValue("eid")
	if pair == "" || eID == "" {
		result := AskBid{
			Success: `false`,
			Message: "Empty Field Selected",
		}
		res, _ := json.Marshal(result)
		fmt.Println(string(res))
		fmt.Fprint(w, string(res))
		return
	}
	if eID == "1" {

		body := bittrex.BittrexMarketData()
		//fmt.Fprint(w, string(body))

		var m interface{}
		err := json.Unmarshal(body, &m)
		if err != nil {
			//panic(err)
			fmt.Println(err)
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
					fmt.Println("Got Sucess As False:", MarketName)

					if pair == MarketName.(string) {

						ask := val2.(map[string]interface{})["Ask"]
						bid := val2.(map[string]interface{})["Bid"]
						//last := val2.(map[string]interface{})["Last"]
						//high24hr := val2.(map[string]interface{})["High"]
						//low24hr := val2.(map[string]interface{})["Low"]
						//vol := val2.(map[string]interface{})["Volume"]
						//baseVol := val2.(map[string]interface{})["BaseVolume"]
						//exchangeID := 2

						result := AskBid{
							Success: `true`,
							Message: "",
							Market:  MarketName.(string),
							Ask:     ask.(float64),
							Bid:     bid.(float64),
						}
						res, _ := json.Marshal(result)
						//fmt.Println(string(res))
						fmt.Fprint(w, string(res))
						return
					}

				}
			}
		}
		result := AskBid{
			Success: `false`,
			Message: "Invalid Pair For The Selected Market",
		}
		res, _ := json.Marshal(result)
		//fmt.Println(string(res))
		fmt.Fprint(w, string(res))
		return
	}

	result := AskBid{
		Success: `false`,
		Message: "Exchange Not Yet Supported",
	}
	res, _ := json.Marshal(result)
	//fmt.Println(string(res))
	fmt.Fprint(w, string(res))
	return

}
