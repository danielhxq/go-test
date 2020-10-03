package dbops

import "database/sql"

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:dbpass!@#@tcp(0.0.0.0:3306)/dev?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}
