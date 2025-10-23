package iteration

import "strings"

func Repeat(character string) string {
	var repeated strings.Builder // stringsBuilder minimizes memory copying

	// Different loops: https://gobyexample.com/for
	for i := 0; i < 5; i++ {
		// repeated += character
		repeated.WriteString(character)
	}

	return repeated.String()
}
