package utils

import (
	"reflect"
	"strconv"
)

func ConvertToInt(v any) int {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Float64:
		return int(v.(float64))
	default: // string
		n, _ := strconv.Atoi(v.(string))
		return n
	}
}
