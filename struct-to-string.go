package toolboxutils

import (
	"fmt"
	"reflect"
	"strings"
)

func StructToString(s interface{}) string {
	value := reflect.ValueOf(s)
	var parts []string

	for i := 0; i < value.NumField(); i++ {
		fieldValue := value.Field(i)
		fieldString := fmt.Sprintf("'%v'", fieldValue.Interface())
		parts = append(parts, fieldString)
	}

	return strings.Join(parts, ", ")
}
