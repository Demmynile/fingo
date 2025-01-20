package db_test

import (
	"database/sql"
	db "github/demmynile/fingo/db/sqlc"
	"log"
	"os"
	"testing"
)

var testQuery *db.Queries

func TestMain(m *testing.M){
		conn , err := sql.Open("postgres", "postgres://spicywords:Harbeedeymee_123@localhost:5432/fingreat_db?sslmode=disable")
		if err != nil{
			log.Fatal("Could not connect to databae" , err)
		}
		testQuery = db.New(conn)

		os.Exit(m.Run())
}