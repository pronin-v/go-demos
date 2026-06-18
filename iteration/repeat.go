package iteration

import "strings"

func Repeat(substr string, iterations int) string {
	var result strings.Builder
	for i := 0; i < iterations; i++ {
		result.WriteString(substr)
	}
	return result.String()
}
