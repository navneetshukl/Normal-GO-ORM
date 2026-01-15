package main

import (
	"log"
	"ngorm"
)

type Table struct {
	ngorm.Table
	Name   string `ngorm:"a,b,c"`
	Mobile string
}

func main() {
	dsn := "host=localhost port=5433 user=postgres password=postgres dbname=postgres sslmode=disable"

	d,err:=ngorm.ConnectToDB(dsn)
	if err!=nil{
		log.Println("Error ",err)
		return
	}

	var t Table
	d.CreateTables(t)
}
