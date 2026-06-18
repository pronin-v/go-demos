package iteration

func Repeat(substr string, iterations int) (result string) {
	for i := 0; i < iterations; i++ {
		result += substr
	}
	return
}
