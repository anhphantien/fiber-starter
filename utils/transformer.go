package utils

import (
	"fmt"
	"reflect"
	"strconv"
)

func ConvertToInt(v any) int {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Float64:
		n, _ := strconv.Atoi(fmt.Sprintf("%.0f", v))
		return n
	default:
		n, _ := strconv.Atoi(v.(string))
		return n
	}
}
