package microservices

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234567"
	dbname   = "itradecoin"
)

type DbCon struct {
	Db *sql.DB
}

var psqlInfo string = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

func OpenConnection() (*DbCon, error) {

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		//panic(err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		fmt.Println("Database connection failed due to", err)
		return nil, err
	}
	fmt.Println("Successfully connected To DB!")

	return &DbCon{
		Db: db,
	}, nil
}

func (con *DbCon) Close() error {
	return con.Db.Close()
}
