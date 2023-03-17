package strings

import "strings"

func MultilineToSingleline(input string) string {
    // Replace all newlines with spaces
    output := strings.ReplaceAll(input, "\n", " ")

    // Remove any extra whitespace
    output = strings.TrimSpace(output)

    return output
}
