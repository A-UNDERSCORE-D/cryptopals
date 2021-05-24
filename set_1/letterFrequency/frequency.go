package letterFrequency

import (
	"strings"
)

// Sourced from https://pi.math.cornell.edu/~mec/2003-2004/cryptography/subs/frequencies.html
var frequenciesText = map[rune]float64{
	'e': 12.02, 't': 9.10, 'a': 8.12, 'o': 7.68, 'i': 7.31, 'n': 6.95, 's': 6.28, 'r': 6.02, 'h': 5.92,
	'd': 04.32, 'l': 3.98, 'u': 2.88, 'c': 2.71, 'm': 2.61, 'f': 2.30, 'y': 2.11, 'w': 2.09, 'g': 2.03,
	'p': 01.82, 'b': 1.49, 'v': 1.11, 'k': 0.69, 'x': 0.17, 'q': 0.11, 'j': 0.1, 'z': 0.07,
	' ': 0.0, // Dont penalise for a space
}

func Frequency(s string) float64 {
	out := 0.0
	for _, v := range strings.ToLower(s) {
		value, exists := frequenciesText[v]
		if exists {
			out += value
		} else {
			out -= 10
		}
	}
	return out
}
