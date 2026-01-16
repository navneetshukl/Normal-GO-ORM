package ngorm

import (
	"reflect"
	"strings"
	"time"
)

// Model this struct is provided by default to create an id,createdAt and updatedAt
type Table struct {
	ID        uint `ngorm:"pk"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func convertToLower(str string) string {
	return strings.ToLower(str)
}

// getColumnType convert go datatype to sql datatype
func getColumnType(ty string) string {
	switch ty {
	case "uint", "int", "uint16", "int16":
		return "INT"
	case "uint32", "int32", "uint64", "int64":
		return "BIGINT"
	case "string":
		return "TEXT"
	case "bool":
		return "BOOLEAN"
	case "Time":
		return "TIMESTAMP"
	}
	return ""
}

// getColumnTag convert the struct tag to valid sql constraint
func getColumnTag(tag string) string {
	tags := strings.Split(tag, ",")
	validTag := ""
	for _, v := range tags {
		switch v {
		case "pk":
			validTag += " PRIMARY KEY"
		}
	}
	return validTag
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
