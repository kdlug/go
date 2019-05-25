package iteration

func main() {

}

func Repeat(char string, count int) string {
	repeated := ""

	for i := 0; i < count; i++ {
		repeated += char
	}
	return repeated
}
