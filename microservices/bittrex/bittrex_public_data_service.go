package bittrex

import (
	"encoding/json"
	"fmt"
	"msservices/microservices"
)

// BittrexMarketData is use to return market data for Bittrex Exchange and will be used as bittrex.BittrexMarketData by other packages;.
func BittrexMarketData() []byte {
	con, err := microservices.OpenConnection()
	if err != nil {
		//return err
		fmt.Println(err)
	}
	defer con.Close()

	body, err := GetTicker("https://bittrex.com/api/v1.1/public/getmarketsummaries")
	//fmt.Println(string(body))
	if err != nil {
		fmt.Println(err)
	}
	var m interface{}
	err = json.Unmarshal(body, &m)
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
				pair := val2.(map[string]interface{})["MarketName"]
				ask := val2.(map[string]interface{})["Ask"]
				bid := val2.(map[string]interface{})["Bid"]
				last := val2.(map[string]interface{})["Last"]
				high24hr := val2.(map[string]interface{})["High"]
				low24hr := val2.(map[string]interface{})["Low"]
				vol := val2.(map[string]interface{})["Volume"]
				baseVol := val2.(map[string]interface{})["BaseVolume"]
				exchangeID := 2

				_, err := con.Db.Exec("INSERT INTO market_data (pair,ask,bid,last,high24hr,low24hr,volume,base_volume,exchange_id)VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)", pair, ask, bid, last, high24hr, low24hr, vol, baseVol, exchangeID)
				if err != nil {
					fmt.Println("Execute Insert Failed Due To: ", err)
				}

			}
		}
	}
	return body
}
