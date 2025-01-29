package main

import (
	"github/demmynile/fingo/api"
)


func main (){
   server := api.NewServer(".")
   server.Start(3000)
}