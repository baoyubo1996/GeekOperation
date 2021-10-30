package week02

import (
	"GeekOperation/src/utils"
	"database/sql"
	"errors"
)

var ErrNoRows = errors.New("<Query> no row found")

func Query(db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query("select  *  from user_go limit 10;")
	utils.CheckErr(err)
	num := 0
	for rows.Next() {
		num++
	}
	if num == 0 {
		return rows, ErrNoRows
	}
	return rows, nil
}
