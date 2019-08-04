package maven

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var pool *sql.DB // Database connection pool.

func connect() {
	// Open up our database connection.
	pool, err := sql.Open("mysql", "root:******@tcp(172.17.0.2:3306)/locrep")

	// if there is an error opening the connection, handle it
	if err != nil {
		fmt.Println("hata aldi")
		panic(err.Error())
	}
	// defer the close till after the main function has finished
	// executing
	defer pool.Close()

	res, _ := pool.Query("SHOW TABLES")
	var table string

	for res.Next() {
		res.Scan(&table)
		fmt.Println(table)
	}
}
