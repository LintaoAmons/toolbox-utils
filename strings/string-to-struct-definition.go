package strings

import "strings"

func StructToStruct(input string) string {
	parts := strings.Split(input, ", ")
	fields := make([]string, len(parts))

	for i, part := range parts {
		fields[i] = strings.Title(part) + " string"
	}

	return "Struct {" + strings.Join(fields, "\n") + "}"
}
