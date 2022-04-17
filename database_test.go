package belajargolangdatabasego

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)


func TestBelajarDatabase(t *testing.T) {
	
}

func TestOpenConnect(t *testing.T){
	db , err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/belajar_golang_database")
	if err != nil {
		panic(err)
	}else {
		fmt.Println("Success Connect")
	}
	defer db.Close()
	//gunakan db

}