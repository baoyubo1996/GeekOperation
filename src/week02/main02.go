package week02

import (
	"GeekOperation/src/utils"
	"fmt"
)

/*
   问题回答：应该 Wrap 这个 error，抛给上层。
   第一版代码
*/
func main() {
	db, errC := Connect()
	utils.CheckErr(errC)
	rows, errR := Query(db)
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
