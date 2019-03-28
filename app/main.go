package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

// TODO
// docs > https://github.com/golang-migrate/migrate/tree/master/database/mysql
// func init() {
// 	db, _ := sql.Open("mysql", "user:pass@tcp(127.0.0.1:3306)/database")
// 	defer db.Close()
// 	driver, _ := mysql.WithInstance(db, &mysql.Config{})
// 	m, _ := migrate.NewWithDatabaseInstance(
// 		"file:///migrations",
// 		"mysql",
// 		driver,
// 	)

// 	m.Steps(2)
// }

func main() {
	db, err := sql.Open("mysql", "user:pass@tcp(127.0.0.1:3306)/database")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))

	//  rows.Scan は引数に `[]interface{}`が必要.
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}
}
