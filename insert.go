package ngorm

import (
	"fmt"
	"reflect"
)

func (d *DBConn) Create(data interface{}) error {

	ty := reflect.TypeOf(data)
	if ty.Kind() == reflect.Ptr {
		ty = ty.Elem()
	}
	if ty.Kind() != reflect.Struct {
		return NotValidStruct
	}
	va:=reflect.ValueOf(data)
	d.generateInsertQuery(ty,va)

	return nil

}

func (d *DBConn) generateInsertQuery(ty reflect.Type,value reflect.Value) {
	//values:=""
	fields:=ty.NumField()
	for i:=0;i<fields;i++{
		fmt.Println("Column Name is ",ty.Field(i).Name)
		fmt.Println("Column Type is ",ty.Field(i).Type.String())
		fmt.Println("Column Value is ",value.Field(i).String())
	}
}
