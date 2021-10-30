package main

import (
	"GeekOperation/src/utils"
	"GeekOperation/src/week02"
	"fmt"
)

/*
   问题回答：应该 Wrap 这个 error，抛给上层。
   第一版代码
*/
func main() {
	db, errC := week02.Connect()
	utils.CheckErr(errC)
	rows, errR := week02.Query(db)
	utils.CheckErr(errR)
	for rows.Next() {
		var id int
		var name int
		var age string
		errS := rows.Scan(&id, &name, &age)
		utils.CheckErr(errS)
		fmt.Println(id, name, age)
	}
}
