package utils

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Model interface {
	TableName() string
}

func GetAllColumnsOfTable(model Model) string {
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
				cases.Title(language.Und).String(model.TableName())+"__"+column,
			))
		}
	}

	if len(s) == 0 {
		return ""
	}
	return ", " + strings.Join(s[:], ", ")
}
