package main

import "ngorm"

type Table struct {
	ngorm.Table
	Name   string `ngorm:"a,b,c"`
	Mobile string
}

func main() {

	var t Table
	ngorm.CreateTables(t)
}
