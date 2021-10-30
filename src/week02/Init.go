package week02

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	return sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/ceramics")
}

func CreteTable(db *sql.DB) (*sql.Stmt, error) {
	return db.Prepare("CREATE TABLE user_go (id int NOT NULL AUTO_INCREMENT, name int(11), age varchar(40), PRIMARY KEY (id));")
}
