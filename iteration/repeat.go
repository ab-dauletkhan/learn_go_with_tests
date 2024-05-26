package iteration

func Repeat(ch string, n int) string {
	var repeated string

	for i := 0; i < n; i++ {
		repeated += ch
	}

	return repeated
}
