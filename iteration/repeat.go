package iteration

const repeatCount = 5

func Repeat(character string) string {
	var repeated string // repeated declared as zero value, for string is "". repeated := ""
	for i := 0; i < repeatCount; i++ {
		repeated += character

	}
	return repeated
}
