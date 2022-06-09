package utils

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"unicode"
)

type Model interface {
	TableName() string
}

func GetAllColumnsOfTableQuery(model Model) string {
	s := []string{}

	r := regexp.MustCompile(`column:\w+`)

	t := reflect.TypeOf(model)
	for i := 0; i < t.NumField(); i++ {
		column := strings.ReplaceAll(
			r.FindString(t.Field(i).Tag.Get("gorm")),
			"column:", "",
		)
		if len(column) > 0 {
			s = append(s, fmt.Sprintf("%v AS %v",
				model.TableName()+"."+column,
				makeFirstLetterUppercase(model.TableName())+"__"+column,
			))
		}
	}

	if len(s) == 0 {
		return ""
	}
	return ", " + strings.Join(s[:], ", ")
}

func makeFirstLetterUppercase(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}
