package ngorm

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
)

type DBConn struct {
	db *sql.DB
}

// ConnectToDB will open connection with DB
func ConnectToDB(connString string) (*DBConn, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("error opening connection with db %s", err)
	}
	return &DBConn{
		db: db,
	}, nil

}

// CreateTables function will create the table by give struct
func CreateTables(table interface{}) error {
	// get column name and column type

	ty := reflect.TypeOf(table)
	if ty.Kind() == reflect.Ptr {
		ty = ty.Elem()
	}
	if ty.Kind() != reflect.Struct {
		return errors.New("provided type is not struct")
	}

	tableName := ty.Name()
	totalFields := ty.NumField()
	fmt.Println("Table Name is ", tableName)
	for i := 0; i < totalFields; i++ {
		field := ty.Field(i)
		fmt.Println("Field is ", field)

		fieldName := field.Name
		fieldType := field.Type
		ormTag := field.Tag.Get("ngorm")

		fmt.Println("Field Name :", fieldName)
		fmt.Println("Field Type :", fieldType)
		fmt.Println("ORM Tag    :", ormTag)
	}
	return nil

}

func parseEmbeddedStructs(ty reflect.Type, fields []*reflect.StructField) {
	if ty.Kind() == reflect.Ptr {
		ty = ty.Elem()
	}
	if ty.Kind() == reflect.Struct {
		parseEmbeddedStructs(ty, fields)
	}
	numFields := ty.NumField()
	for i := 0; i < numFields; i++ {

	}

}
