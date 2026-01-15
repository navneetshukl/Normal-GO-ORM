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

	t := reflect.TypeOf(table)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return errors.New("provided type is not struct")
	}
	fmt.Println("Table:", t.Name())
	fields := flattenStruct(t)
	for _, f := range fields {
		fmt.Println("Field:", f.Name, "Type:", f.Type)
	}

	return nil
}

// flattenStruct function should flatten the struct and embedded struct
func flattenStruct(t reflect.Type) []reflect.StructField {
	fields := []reflect.StructField{}

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return fields
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		ft := f.Type

		if ft.Kind() == reflect.Ptr {
			ft = ft.Elem()
		}
		if shouldFlatten(ft) {
			allFields := flattenStruct(ft)
			fields = append(fields, allFields...)
			continue
		}
		fields = append(fields, f)
	}

	return fields
}

func shouldFlatten(ft reflect.Type) bool {
	return ft.Kind() == reflect.Struct &&
		ft.PkgPath() != "" &&
		ft.PkgPath() != "time" &&
		ft.PkgPath() != "database/sql"
}
