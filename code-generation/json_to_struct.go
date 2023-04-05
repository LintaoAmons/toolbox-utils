package code_generation

import (
	"encoding/json"
	"fmt"
	"strings"
)

func JsonToStructInstance(varName,typeName string, jsonStr string) (string, error) {
	var m map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		return "", err
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("var %s = %s{", varName, typeName))

	for k, v := range m {
		sb.WriteString(fmt.Sprintf(" %s: ", k))
		switch val := v.(type) {
		case string:
			sb.WriteString(fmt.Sprintf("\"%s\",", val))
		case int:
			sb.WriteString(fmt.Sprintf("%d,", val))
		case float64:
			sb.WriteString(fmt.Sprintf("%f,", val))
		case bool:
			sb.WriteString(fmt.Sprintf("%t,", val))
		default:
			sb.WriteString(fmt.Sprintf("%v,", val))
		}
	}

	sb.WriteString(" }")

	return sb.String(), nil
}
