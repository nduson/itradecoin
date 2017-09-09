package bittrex

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetTicker(url string) (bs []byte, err error) {
	fmt.Println("Enter Get Ticker Function: Getting Ticker From URL: " + url + "")
	//url := "https://poloniex.com/public?command=returnTicker"
	res, err := http.Get(url)
	if (err) != nil {
		//fmt.Println("ERROR: Failed To Connected to " + url + " For Market Data")
		return nil, err
	}
	defer res.Body.Close()
	bs, err = ioutil.ReadAll(res.Body)
	if (err) != nil {
		//panic(err)
		return nil, err
	}
	//fmt.Println(string(body))
	res.Body.Close()
	return bs, nil
}
