package ngorm

import "errors"

var (
	NotValidStruct error = errors.New("provided type is not struct")
)
