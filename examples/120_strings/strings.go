package strings

func Reverse(s string) string {
	// strings are read-only so we should use bytes or rune
	// bytes are fixed length, so it won't work properly
	// to represent any character it's better to use rube
	b := []rune(s)

	// abcde
	// 1. we can divide string to 2, if string is 5 charactrers it will be ab|c|de
	// 2. we switch first and last element in 1st iteration eb|c|da
	// 3. In 2nd iteration we take second and (len-1-i) ed|c|ba
	// 4. Element in the middle will stay in it's position

	for i := 0; i < len(b)/2; i++ {
		j := len(b) - 1 - i // take 1 element from the end
		// and switch characters
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}
