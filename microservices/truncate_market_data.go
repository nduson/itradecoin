package microservices

import (
	"fmt"
)

// TruncateMarketData this delect market data after set time interval
func TruncateMarketData() {
	con, err := OpenConnection()
	if err != nil {
		//return err
		fmt.Println(err)
	}
	defer con.Close()
	fmt.Println("Enter truncate market data sucessfully")
	//for {
	fmt.Println("truncateMarketData Func waiting for 24hrs to truncate data")

	//time.Sleep(100 * time.Second)

	res, err := con.Db.Exec("DELETE FROM market_data WHERE ctid IN (SELECT ctid FROM market_data LIMIT 100)")
	if err != nil {
		fmt.Println("Execute Insert Failed Due To: ", err)
	}
	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println("Row Affected Failed Due To: ", err)
	}

	fmt.Println("Sucessfully truncated market data; No Of Rows Affected: ", rowCount)
	//}
}
