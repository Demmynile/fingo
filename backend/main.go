package main

import (
	"github.com/demmynile/fingo/api"
)


func main (){
   server := api.NewServer(".")
   server.Start(8000)
}