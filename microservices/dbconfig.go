package microservices

import (
	"database/sql"
	"fmt"
	// this imports postgress drivers for db.
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234567"
	dbname   = "itradecoin"
)

// DbCon this hold pointer to db
type DbCon struct {
	Db *sql.DB
}

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

// OpenConnection is use open the connection
func OpenConnection() (*DbCon, error) {

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		//panic(err)
		//return nil, err
		fmt.Println("Database connection failed due to", err)
	}

	if err = db.Ping(); err != nil {
		fmt.Println("Database connection failed due to", err)
		//return nil, err
	}
	fmt.Println("Successfully connected To DB!")

	return &DbCon{
		Db: db,
	}, nil
}

// Close  our db conncetion.
func (con *DbCon) Close() error {
	return con.Db.Close()
}

//qy := "INSERT INTO market_data (pair,ask,bid,last,high24hr,low24hr,volume,base_volume,exchange_id)VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)"
//stmt, _ := con.Db.Prepare(qy)
/*if err != nil {
	fmt.Println("Insert Failed Due To: ", err)
	//fmt.Println(err)

}*/
//_, err = stmt.Exec(pair, ask, bid, last, high24hr, low24hr, vol, base_vol, exchange_id)
//"INSERT INTO market_data (pair,ask,bid,last,high24hr,low24hr,volume,base_volume,exchange_id)VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)"
//checkErr(err)
//_, err := con.Db.Exec("INSERT INTO market_data (pair,ask,bid,last,high24hr,low24hr,volume,base_volume,exchange_id)VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)", pair, ask, bid, last, high24hr, low24hr, vol, base_vol, exchange_id)
//if err != nil {
//fmt.Println("Execute Insert Failed Due To: ", err)
//fmt.Println(err)
//}
