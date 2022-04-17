package belajargolangdatabasego

import (
	"database/sql"
	"fmt"
	"time"
)

func GetConnection() *sql.DB{
	db , err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/belajar_golang_database?parseTime=true")
	if err != nil {
		panic(err)
	}else {
		fmt.Println("Success Connect")
	}

	db.SetMaxIdleConns(10) //jumlah minimal koneksi yang akan dibuat
	db.SetMaxOpenConns(100) //jumlah maksimal koneksi
	db.SetConnMaxIdleTime(5 * time.Minute) //lamanya waktu untuk koneksi yang sudah tidak digunakan akan dihapus
	db.SetConnMaxLifetime(60 * time.Minute) //lamanya koneksi dapat digunakan
	
	return db
}