package db

import (
	"log"
	"testing"
)

func TestMain(m *testing.M){
		conn , err := sql.open("postgres", "postgres://spicywords:Harbeedeymee_123@localhost:5432/fingreat_db?sslmode=disable")
		if err != nil{
			log.Fatal("Could not connect to databae" , err)
		}
}