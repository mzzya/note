package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

type Test struct {
	ID   int `gorm:"primary_key"`
	Name string
}

func (t Test) TableName() string {
	return "test1"
}

func main() {
	// conn, err := sql.Open("mssql", "sqlserver://sa:123123@192.168.1.3:1433?database=test")
	// if err != nil {
	// 	log.Fatal("Open connection failed:", err.Error())
	// }
	// defer conn.Close()

	// result, err := conn.Query("insert into test1 values('hello12346');select ID = convert(bigint, SCOPE_IDENTITY())")
	// if err != nil {
	// 	fmt.Printf("insert into error:%s\n", err)
	// 	return
	// }
	// // LastInsertId, err := result.LastInsertId()
	// // if err != nil {
	// // 	fmt.Printf("LastInsertId:%s\n", err)
	// // 	return
	// // }
	// // fmt.Println("LastInsertId", LastInsertId)
	// // RowsAffected, err := result.RowsAffected()
	// // if err != nil {
	// // 	fmt.Printf("RowsAffected:%s\n", err)
	// // 	return
	// // }
	// // fmt.Println("RowsAffected", RowsAffected)
	// fmt.Printf("%#v\n", result)
	db, err := gorm.Open("mssql", "sqlserver://sa:123123@192.168.1.3:1433?database=test")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	t := &Test{Name: "aaak11lk1"}
	rs := db.Create(t).
	if rs != nil {
		fmt.Printf("rs.Error:%s\n", rs)
	}
	fmt.Println("ID", t.ID)
}
