package belajargolangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestConnectionDatabase(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/belajar")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

func TestExecuteSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	context := context.Background()

	script := "insert into customer(id, name) values('rizki1', 'Rizki');"
	_, err := db.ExecContext(context, script)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success Insert to database")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	context := context.Background()

	script := "select id, name from customer"
	rows, err := db.QueryContext(context, script)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("id", id)
		fmt.Println("name", name)
	}
	defer rows.Close()
}
