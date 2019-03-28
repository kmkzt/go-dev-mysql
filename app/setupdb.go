package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

// docs > https://github.com/golang-migrate/migrate/tree/master/database/mysql
func setupdb() {
	db, _ := sql.Open("mysql", "user:pass@tcp(127.0.0.1:3306)/database?multiStatements=true")
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"mysql",
		driver,
	)

	m.Steps(2)
}
