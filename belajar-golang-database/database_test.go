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

func TestQuerySql2(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	context := context.Background()

	script := "select id, name, email, balance, rating, birth_date, married, created_at from customer"
	rows, err := db.QueryContext(context, script)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id, name, email sql.NullString
		var balance int32
		var rating float64
		var birthDate, createAt sql.NullTime
		var married bool
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("=============================")
		fmt.Println("id", id)
		fmt.Println("name", name)
		if email.Valid {
			fmt.Println("email", email.String)
		}
		fmt.Println("balance", balance)
		fmt.Println("rating", rating)
		if birthDate.Valid {
			fmt.Println("birthDate", birthDate.Time)
		}
		fmt.Println("createAt", createAt)
		fmt.Println("married", married)
		fmt.Println("=============================")
	}
	defer rows.Close()
}
