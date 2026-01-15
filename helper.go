package ngorm

import (
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
