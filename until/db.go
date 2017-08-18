package until

import (
	"database/sql"
	// it is for mysql
	_ "github.com/go-sql-driver/mysql"
)

// DB function
func DB() *sql.DB {
	var path string;
	path = GetElement("dbuser")+":"+GetElement("password")+"@tcp("+GetElement("host")+")/"+GetElement("dbname")
	db, _ := sql.Open("mysql", path)
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
