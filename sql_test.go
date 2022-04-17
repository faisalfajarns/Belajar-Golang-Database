package belajargolangdatabasego

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

type Data []string

//notes untuk tipe data sql yang bisa null diharapkan menggunakan tipe data sql.nullable 
type Results struct{
	id string 
	name string
	email sql.NullString
	balance int32
	rating float64
	birthDate sql.NullTime
	married bool
}
func TestExecInsertSql(t *testing.T) {
	db := GetConnection()
	
	defer db.Close()
	
	ctx := context.Background()
	
	bytes := make([]byte, 12)
	_, newerr := rand.Read(bytes)
	if newerr !=nil {
		panic(newerr)
	}

	id := hex.EncodeToString(bytes)
	fmt.Println("bytes",len(id))

	
	//cara 1
	// queries := "INSERT INTO customers(id, name) VALUES('C0001', 'Faisal')" 
	// _, err := db.ExecContext(ctx, queries) //exec contex digunakan untuk sql yg tidak membutuhkan hasil data seperti update delete insert
	
	//cara 2 dengan paramter ke 3

	utc := time.Date(1997, 10, 25, 10, 10, 10, 10, time.UTC)

	data :=  []interface{}{
		id,
		strings.ToLower("Fajar"),
		"fajar@gmail.com",
		200000,
		4.0,
		utc,
		false,
	}

	params := &Data{
		"id",
		"name",
		"email",
		"balance",
		"rating",
		"birth_date",
		"married",
	}

	formatted := strings.Join(*params,", ")

	queries := fmt.Sprintf("INSERT INTO customers(%s) VALUES(?, ?, ?, ?, ?, ?, ?)", formatted) 
	//note params ke 3 untuk parameter query
		_, err:= db.ExecContext(ctx, queries, data...)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert New Customer")
}

func TestExecDeleteSql(t *testing.T) {
	db := GetConnection()
	
	defer db.Close()
	
	ctx := context.Background()

	queries := "DELETE from customers where id = '1365c914227f3629539c3668'" 
	_, err := db.ExecContext(ctx, queries)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success Delete Customer")
}

func TestExecUpdateSql(t *testing.T) {
	db := GetConnection()
	
	defer db.Close()
	
	ctx := context.Background()

	queries := "UPDATE customers SET name = 'Nursaid' WHERE id = 'C0001'" 
	_, err := db.ExecContext(ctx, queries)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success Update Customer")
}

func TestSprintF(t *testing.T) {
	data := &Data{
		"id",
		"name",
	}

	formatted := strings.Join(*data,", ")
	// fmt.Println()
	query := fmt.Sprintf("INSERT INTO customers(%s) VALUES(?,?)", formatted)

	fmt.Println(query)
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	
	defer db.Close()
	
	ctx := context.Background()

	// data :=  []interface{}{
	// 	"7910f5c74442b86fa94b0358",
	// }

	params := &Data{
		"id",
		"name",
	}

	formatted := strings.Join(*params, ", ")

	queries := fmt.Sprintf("SELECT %s FROM customers", formatted)

	rows, err:= db.QueryContext(ctx, queries)
	if err != nil {
		panic(err)
	}

	res := []Results{}

	for rows.Next(){
		each := Results{}
		// var id, name string
		err := rows.Scan(&each.id, &each.name)
		if err != nil {
			panic(err)
		}

		res = append(res, each)
		// fmt.Println("data::", res)


		// fmt.Println("Name:", name)
	}
	
	for _, each := range res{
	fmt.Println("Name:", each.name)
	}

	defer rows.Close()
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	
	defer db.Close()
	
	ctx := context.Background()

	// data :=  []interface{}{
	// 	"7910f5c74442b86fa94b0358",
	// }

	params := &Data{
		"id",
		"name",
		"email",
		"balance",
		"rating",
		"birth_date",
		"married",
	}

	formatted := strings.Join(*params, ", ")

	queries := fmt.Sprintf("SELECT %s FROM customers", formatted)

	rows, err:= db.QueryContext(ctx, queries)
	if err != nil {
		panic(err)
	}

	res := []Results{}

	for rows.Next(){
		each := Results{}
		// var id, name string
		err := rows.Scan(&each.id, &each.name, &each.email, &each.balance, &each.rating, &each.birthDate, &each.married)
		if err != nil {
			panic(err)
		}

		res = append(res, each)
		// fmt.Println("data::", res)


		// fmt.Println("Name:", name)
	}
	
	for _, each := range res{
	fmt.Println("==================")
	fmt.Println("Name:", each.name)
	if each.email.Valid {
		fmt.Println("Email:", each.email.String)
	}
	fmt.Println("Balance:", each.balance)
	fmt.Println("Rating:", each.rating)
	if each.birthDate.Valid {
		fmt.Println("Birth:", each.birthDate.Time)
	}
	fmt.Println("Married:", each.married)
	}

	defer rows.Close()
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	
	defer db.Close()
	
	ctx := context.Background()

	username := "admin'; #"
	password := "salah"

	queries := "SELECT username FROM users WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
	fmt.Println(queries)
	rows, err:= db.QueryContext(ctx, queries)
	if err != nil {
		panic(err)
	}

	if rows.Next(){
		var username string
		err := rows.Scan(&username)
			if err != nil {
				panic(err)
			}
		fmt.Println("Sukses Login")
	}else {
		fmt.Println("Gagal Login")
	}

}

func TestSqlWithParams(t *testing.T) {
	db := GetConnection()
	
	defer db.Close()
	
	ctx := context.Background()

	username := "admin"
	password := "admin"

	data := []interface{}{
		username,
		password,
	}

	fmt.Println(data...)
	 

	queries := "SELECT username FROM users WHERE username = ? AND password = ? LIMIT 1"
	rows, err:= db.QueryContext(ctx, queries, data...)
	if err != nil {
		panic(err)
	}

	if rows.Next(){
		var username string
		err := rows.Scan(&username)
			if err != nil {
				panic(err)
			}
		fmt.Println("Sukses Login")
	}else {
		fmt.Println("Gagal Login")
	}

}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	
	defer db.Close()
	
	ctx := context.Background()
	
	email := "faisal@gmail.com"
	comment := "test komen"
	queries := "INSERT INTO comments(email, comment) VALUES(?, ?)"

	res, err := db.ExecContext(ctx, queries, email, comment)
	if err != nil {
		panic(err)
	}
	// notes last insert id digunakan untuk mendapatkan id jika idnya berupa increment
	insertId, err := res.LastInsertId()
		if err != nil {
			panic(err)
		}

	fmt.Println("Sukses :", insertId )
}

func TestPrepareStatement(t *testing.T) {
	//notes prepare statement digunakan untuk query yang sama 
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	// comment := "test 2 komen prepare"
	queries := "INSERT INTO comments(email, comment) VALUES(?, ?)"

	stmt, err := db.PrepareContext(ctx, queries)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := "Faisal"+strconv.Itoa(i)+"@gmail.com"
		comment := "Test Komen "+ strconv.Itoa(i+1)
		
		res, err := stmt.ExecContext(ctx, email, comment)

		if err != nil {
			panic(err)
		}

		id, err := res.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id", id)
	}

}

func TestTransaction(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()
	
	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}
	// do transaction here
	queries := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	for i := 0; i < 10; i++ {
		email := "Faisal"+strconv.Itoa(i)+"@gmail.com"
		comment := "Test Komen "+ strconv.Itoa(i+1)
		
		res, err := tx.ExecContext(ctx, queries,email, comment)

		if err != nil {
			panic(err)
		}

		id, err := res.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id", id)
	}
	err = tx.Rollback()

	if err != nil {
		panic(err)
	}

}