package belajargolangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
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

func TestSQLInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "rizki'; #"
	password := "rizki1"
	script := "select username from user where username = '" + username + "' and password = '" + password + "' LIMIT 1"
	fmt.Println(script)
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("sukses login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestSQLQueryParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "rizki'; #"
	password := "rizki1"
	script := "select username from user where username = ? and password = ? LIMIT 1"
	fmt.Println(script)
	rows, err := db.QueryContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("sukses login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestSQLExecuteParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	context := context.Background()

	username := "admin"
	password := "password"
	script := "insert into user(username, password) values(?, ?)"
	_, err := db.ExecContext(context, script, username, password)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success Insert to database")
}

func TestSQLExecuteAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	context := context.Background()

	email := "rizki@gmail.com"
	comment := "ini adalah contoh comment"
	script := "insert into comments(email, comment) values(?, ?)"
	result, err := db.ExecContext(context, script, email, comment)
	if err != nil {
		panic(err)
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("Success Insert to database dengan id :", lastId)
}

func TestSQLExecutePreparedStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	context := context.Background()

	script := "insert into comments(email, comment) values(?, ?)"
	statement, err := db.PrepareContext(context, script)
	if err != nil {
		panic(err)
	}

	for i := 1; i <= 10; i++ {
		email := "rizki" + strconv.Itoa(i) + "@gmail.com"
		comment := "ini adalah contoh comment ke " + strconv.Itoa(i)

		result, err := statement.ExecContext(context, email, comment)
		if err != nil {
			panic(err)
		}
		lastId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Berhasil Insert dengan id", lastId)
	}

	defer statement.Close()
}

func TestSQLQueryPreparedStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	context := context.Background()

	script := "select email, comment from comments"
	statement, err := db.PrepareContext(context, script)
	if err != nil {
		panic(err)
	}

	result, err := statement.QueryContext(context)
	if err != nil {
		panic(err)
	}

	for result.Next() {
		var email, comment string
		err := result.Scan(&email, &comment)
		if err != nil {
			panic(err)
		}
		fmt.Println("email", email)
		fmt.Println("comment", comment)
	}

	defer statement.Close()
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	script := "insert into comments(email, comment) values(?, ?)"
	for i := 1; i <= 10; i++ {
		email := "rizki" + strconv.Itoa(i) + "@gmail.com"
		comment := "ini adalah contoh comment ke " + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil {
			tx.Rollback()
			panic(err)
		}
		lastId, err := result.LastInsertId()
		if err != nil {
			tx.Rollback()
			panic(err)
		}
		fmt.Println("Berhasil Insert dengan id", lastId)
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}
