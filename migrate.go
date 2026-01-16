package ngorm

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"

	_ "github.com/lib/pq"
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
func (d *DBConn) CreateTables(table interface{}) error {

	t := reflect.TypeOf(table)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return NotValidStruct
	}
	//fmt.Println("Table:", t.Name())
	fields := flattenStruct(t)
	// for _, f := range fields {
	// 	fmt.Println("Field:", f.Name, "Type:", f.Type, "Tag:", f.Tag.Get("ngorm"))
	// }
	sql := d.generateSQL(t.Name(), fields)
	fmt.Println("SQL ", sql)

	_, err := d.db.ExecContext(context.Background(), sql)
	if err != nil {
		return fmt.Errorf("error in creating the table %s ", err)
	}

	return nil
}

// generateSQL will generate the valid sql to create table
func (d *DBConn) generateSQL(tableName string, fields []reflect.StructField) string {
	columns := ""
	for _, f := range fields {
		columnName := convertToLower(f.Name)
		columnType := getColumnType(f.Type.Name())
		columnTags := getColumnTag(f.Tag.Get("ngorm"))
		if columns == "" {
			columns = fmt.Sprintf("%s %s%s, \n", columnName, columnType, columnTags)
		} else {
			columns = fmt.Sprintf("%s %s %s%s, \n", columns, columnName, columnType, columnTags)
		}
	}
	if len(columns) >= 3 {
		columns = columns[:len(columns)-3]
	}
	tableName = convertToLower(tableName)
	if tableName[len(tableName)-1] != 's' {
		tableName = fmt.Sprintf("%ss", tableName)
	}
	columns = fmt.Sprintf("(\n%s\n);", columns)
	sqlCommand := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s%s", tableName, columns)
	return sqlCommand

}
