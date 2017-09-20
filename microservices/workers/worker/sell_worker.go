package worker

import (
	"encoding/json"
	"fmt"
	"math"
	"msservices/microservices"
)

// SellWorker is func to update buy order
func SellWorker() {

	fmt.Println("Entered Our buyOrderUpdateWorker func To Get The Requestd Data From DB")
	con, err := microservices.OpenConnection()
	if err != nil {
		//return err
		fmt.Println(err)
	}
	defer con.Close()

	row, err := con.Db.Query("SELECT id_sell_worker,market,exchange_id,highest_bid_price,lowest_bid_price,highest_ask_price,lowest_ask_price,highest_volume,lowest_volume,actual_rate,actual_quantity,profit_keep,sell_trigger,last_volume,start_volume FROM sell_worker WHERE work_status = 0")
	if err != nil {
		fmt.Println("Select Failed Due To: ", err)
	}
	defer row.Close()

	for row.Next() {
		fmt.Println("Entered row dot Next")
		var workerID, market, exchangeID string
		var highBid, lowBid, highAsk, lowAsk, highVol, lowVol, lastVol, startVol, actualRate, actualQty, profitKeep, selTrigger float64
		err = row.Scan(&workerID, &market, &exchangeID, &highBid, &lowBid, &highAsk, &lowAsk, &highVol, &lowVol, &actualRate, &actualQty, &profitKeep, &selTrigger, &lastVol, &startVol)
		if err != nil {
			fmt.Println("Row Scan Failed Due To: ", err)
		}
		//http://localhost:5000/pair/price?pair=btc-bcc&eid=1
		// call the end point with the gotten values.
		body, err := microservices.GetTicker("http://localhost:5000/pair/price?pair=" + market + "&eid=" + exchangeID + "")
		fmt.Println(string(body))
		if err != nil {
			fmt.Println("Error On Bittrex GetTicker Func", err)
			return
		}
		// unmarshal the json response.
		var m map[string]interface{}
		err = json.Unmarshal(body, &m)
		if err != nil {
			//panic(err)
			fmt.Println(err)
		}

		//pair := m["market"]
		ask := m["ask"]
		bid := m["bid"]
		//high := m["high"]
		//low := m["low"]
		vol := m["volume"]
		fmt.Println("Values first time:", lowAsk, lowBid, lowVol)
		/// First time insert into the sell worker table.
		if lowAsk == 0 && lowBid == 0 && lowVol == 0 {
			fmt.Println("Values first time:", lowAsk, lowBid, lowVol)
			fmt.Println("First time insert into the sell worker table.")
			_, err = con.Db.Exec("UPDATE sell_worker SET highest_bid_price = $1, lowest_bid_price= $2,highest_ask_price =$3,lowest_ask_price=$4,start_volume=$5,highest_volume=$6,lowest_volume=$7,last_volume=$8 WHERE id_sell_worker = $9", bid, bid, ask, ask, vol, vol, vol, vol, workerID)
			if err != nil {
				fmt.Println("Update Failed Due To: ", err)
			}
			return
		}

		/// check if high ask price have changed
		if ask.(float64) > highAsk {
			highAsk = ask.(float64)
		}
		fmt.Println("HighestAsk: ", highAsk)

		/// check if low ask price have changed
		if ask.(float64) < lowAsk {
			lowAsk = ask.(float64)
		}
		fmt.Println("LowestAsk: ", lowAsk)

		/// check if high bid price have changed
		if bid.(float64) > highBid {
			highBid = bid.(float64)
		}
		fmt.Println("HighestBid: ", highBid)

		/// check if low bid price have changed
		if bid.(float64) < lowBid {
			lowBid = bid.(float64)
		}
		fmt.Println("LowestBid: ", lowBid)
		/// check if high bid price have changed
		if bid.(float64) > highVol {
			highVol = vol.(float64)
		}
		fmt.Println("HighestVol: ", highVol)

		/// check if low bid price have changed
		if bid.(float64) < lowVol {
			lowVol = vol.(float64)
		}
		fmt.Println("Lowestvol: ", lowVol)

		// check vol_diff
		// volume_diff (last_volume - start_volume)
		volDiff := RoundDown((lastVol - startVol), 8)
		fmt.Println("Volume Difference: ", volDiff)

		// check vol percent
		// vol_percent (volume_diff/start_volume).
		volPercent := RoundDown(((volDiff / startVol) * 100), 4)
		fmt.Println("Volume Percent: ", volPercent)

		/// check threshold
		//⁠⁠⁠threshold = ActualRate + [(HighestBid - ActualRate) × (profit_keep + sell_trigger)/100]
		threshold := RoundDown(actualRate+RoundDown(((highBid-actualRate)*RoundDown((profitKeep+selTrigger), 8)/100), 8), 8)
		fmt.Println("threshold: ", threshold)

		// check exit price
		// But the exit price = (profit_keep + (sell_trigger÷2))is evaluated between entry_price and highest price.
		exitPrice := RoundDown((profitKeep + RoundDown((selTrigger/2), 8)), 8)
		fmt.Println("Exit Price: ", exitPrice)

		// check sell_price
		//ActualRate + [(HighestBid - ActualRate) × (profit_keep + (sell_trigger÷2))/100]
		sellPrice := RoundDown(actualRate+RoundDown(((highBid-actualRate)*RoundDown((profitKeep+(selTrigger/2)/100), 8)), 8), 8)
		fmt.Println("SellPrice: ", sellPrice)

		// check high profit_keep
		// high_profit: which is d profit in terms of [(current bid - actual_rate) - 0.25%)]
		highProfit := RoundDown(((bid.(float64) - actualRate) - (0.25 / 100)), 8)
		fmt.Println("high profit: ", highProfit)

		// check high profit percentage
		// high_profit_perc: (high_profit÷actual_rate × 100).
		highProfitPercent := RoundDown(((highProfit / actualRate) * 100), 4)
		fmt.Println("High Profit Percent: ", highProfitPercent)

		// check for actual_locked_profit
		// actual_locked_profit: [(sell_price - actual_rate) - 0.25%].
		actualLockProfit := RoundDown(((sellPrice - actualRate) - (0.25 / 100)), 8)
		fmt.Println("Actual Lock Profit: ", actualLockProfit)

		// check actual_locked_perc_profit
		// actual_locked_perc_profit : (actual_locked_profit ÷ actual_rate × 100)
		actualLockProfitPer := RoundDown(((actualLockProfit / actualRate) * 100), 4)
		fmt.Println("Actual Lock Profit Percent: ", actualLockProfitPer)

		// check locked_proceed
		// locked_proceed: which is PCw = (SCw x sell_price) - 0.25% translating to (actual_quantity × sell_price) - 0.25%
		lockedProceed := RoundDown(((actualQty * sellPrice) - (0.25 / 100)), 8)
		fmt.Println("Locked Procced: ", lockedProceed)

		// check last_proceed
		// last_proceed: It is to b computed from [(actual_quantity × last_bid) - 0.25%]
		lastProceed := RoundDown((actualQty*bid.(float64))-(0.25/100), 8)
		fmt.Println("Last Procced: ", lastProceed)
		// check PL
		//  PL which tells the current profit/last profit is to be computer from [(last_bid - actual_rate) - 0.25%]
		pL := RoundDown((bid.(float64)-actualRate)-(0.25/100), 8)
		fmt.Println("PL: ", pL)
		// check per pL
		// percent_profit which is expression of last/current profit is to b computed = (PL ÷ actual_rate × 100).
		percentPL := RoundDown(((pL / actualRate) * 100), 4)
		fmt.Println("PL Percent: ", percentPL)

		_, err = con.Db.Exec("UPDATE sell_worker SET highest_bid_price = $1, lowest_bid_price= $2,highest_ask_price =$3,lowest_ask_price=$4,last_bid=$5,highest_volume=$6,lowest_volume=$7,threshold=$8,sell_price=$9,high_profit=$10,high_profit_perc=$11,actual_locked_profit=$12,actual_locked_perc_profit=$13,locked_proceed=$14,last_proceed=$15,pl=$16,exit_price=$17,last_volume=$18,volume_diff=$19,vol_percent=$20 WHERE id_sell_worker = $21",
			highBid, lowBid, highAsk, lowAsk, bid, highVol, lowVol, threshold, sellPrice, highProfit, highProfitPercent, actualLockProfit,
			actualLockProfitPer, lockedProceed, lastProceed, pL, exitPrice, vol, volDiff, volPercent, workerID)
		if err != nil {
			fmt.Println("Update Failed Due To: ", err)
		}
	}

}

// Round var f float64 = 514.89317306
// Round this rounds Output: 514.89317306 to 514
func Round(input float64) float64 {
	if input < 0 {
		return math.Ceil(input - 0.5)
	}
	return math.Floor(input + 0.5)
}

// RoundUp sample output in 4 decmial places
// var f float64 = 514.89317306
//RoundUp this rounds Output: 514.89317306 to 514.8932
func RoundUp(input float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * input
	round = math.Ceil(digit)
	newVal = round / pow
	return
}

// RoundDown Output in 4 decmial places
// var f float64 = 514.89317306
//RoundDown this round Output: 514.89317306 to 514.8931
func RoundDown(input float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * input
	round = math.Floor(digit)
	newVal = round / pow
	return
}
