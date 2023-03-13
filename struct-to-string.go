package toolboxutils

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

func StructToString(s interface{}) string {
	// Get the reflect.Value of the struct
	value := reflect.ValueOf(s)

	var parts []string

	// Loop over the fields of the struct
	for i := 0; i < value.NumField(); i++ {
		// Get the reflect.Value of the field
		fieldValue := value.Field(i)

		// Check if the field type is time.Time
		if fieldValue.Type() == reflect.TypeOf(time.Time{}) {
			// Format the time.Time field as a string
			fieldString := fmt.Sprintf("'%s'::timestamp WITH TIME ZONE", fieldValue.Interface().(time.Time).Format("2006-01-02 15:04:05.999999 -07:00"))
			parts = append(parts, fieldString)
		} else {
			// Convert the field value to a string and add it to the parts slice
			fieldString := fmt.Sprintf("'%v'", fieldValue.Interface())
			parts = append(parts, fieldString)
		}
	}

	// Join the parts slice into a comma-separated string
	return strings.Join(parts, ", ")
}
