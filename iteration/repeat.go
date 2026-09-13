package iteration

func Repeat(character string) string {
	var repeated string // repeated declared as zero value, for string is "". repeated := ""
	for i := 0; i < 5; i++ {
		repeated += character

	}
	return repeated
}

