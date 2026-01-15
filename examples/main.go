package main

import "ngorm"

type Table struct {
	ngorm.Model
	Name   string `ngorm:"a,b,c"`
	Mobile string
}

func main() {

	var t Table

	ngorm.CreateTables(t)
}
