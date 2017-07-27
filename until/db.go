package until

import (
	"database/sql"
	// it is for mysql
	_ "github.com/go-sql-driver/mysql"
)

// DB function
func DB() *sql.DB {
	db, _ := sql.Open("mysql", "root:123456@/sample")
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
