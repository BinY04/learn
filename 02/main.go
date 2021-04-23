package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	errors2 "github.com/pkg/errors"
	"log"
)

func main() {
	db, err := sql.Open("mysql",
		"root:mysqldata@tcp(127.0.0.1:3306)/goodbuy")
	if err != nil {
		log.Fatal(err)
	}

	querySql := "select username from gb_admin where id = ?"
	var data string
	data, err = Select(db, querySql, 100)
	if err != nil {
		log.Fatal(fmt.Printf("err :%+v\n", err))
	}
	log.Print(data)
	defer db.Close()
}

func Select(db *sql.DB, querySql string, param int) (string, error) {
	var name string
	err := db.QueryRow(querySql, param).Scan(&name)
	if err != nil {
		// 如果是查询没有数据，应该只记录日志
		if errors.Is(err, sql.ErrNoRows) {
			log.Fatalf("sql:%s param: %v; query not data", querySql, param)
		} else {
			return "", errors2.Wrap(err, "query error")
		}
	}
	return name, err
}
