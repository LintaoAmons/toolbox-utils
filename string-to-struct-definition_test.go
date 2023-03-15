package toolboxutils_test

import (
	"fmt"
	"testing"

	toolboxutils "github.com/LintaoAmons/toolbox-utils"
	"github.com/atotto/clipboard"
)

func TestStringToStructDef(t *testing.T) {
	input := "id, id_number, application_number, status, created_at, created_by, application_date, confirmation_date, marriage_date, marriage_source, marriage_cert, allow_change, remarks, marriage_type, marriage_status, age_type, group_type, lodgement_of_caveat, is_re_registration_application_viewed, service_status, data_source, allow_online_declaration, declaration_status, straight_through_status"
	expected := "Struct {Field1: string, Field2: string, Field3: string}"

	result := toolboxutils.StructToStruct(input)
	fmt.Printf(result)
	clipboard.WriteAll(result)

	if result != expected {
		t.Errorf("stringToStructDef(%q) returned %q, expected %q", input, result, expected)
	}
}
