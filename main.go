package main

import (
	"ACAT/model"
	"ACAT/routes"
	_ "github.com/go-sql-driver/mysql"
)

func main(){

	model.IntDb()
	routes.InitRouter()

}

