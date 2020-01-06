package main

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	sql.Open("sqlserver", "")
}
