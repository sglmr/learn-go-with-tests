package iteration

const RepeatCount = 5

func Repeat(character string) string {
	var repeated string

	for i := 0; i < RepeatCount; i++ {
		repeated = repeated + character
	}
	return repeated
}

func ExampleRepeat(character string, count int) string {
	var repeated string

	for i := 0; i < count; i++ {
		repeated = repeated + character
	}
	return repeated
}
