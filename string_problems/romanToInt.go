package string_problems

func romanToInt(s string) int {
	numbers := make([]int, len(s))

	// Get string
	for i, letter := range s {
		numbers[i] = singularRomanDict[string(letter)]
	}

	// replace with values
	sum := 0

	for i := 0; i < len(numbers); i++ {
		if i > 0 {
			if numbers[i] > numbers[i-1] {
				sum += numbers[i] - numbers[i-1]*2
				continue
			}
		}
		sum += numbers[i]
	}
	// if the value next to it is higher than the current one subtract

	return sum
}

type RomanDictionary map[string]int

var singularRomanDict = RomanDictionary{
	"I": 1,
	//"IV": 4,
	"V": 5,
	//"IX": 9,
	"X": 10,
	//"XL": 40,
	"L": 50,
	//"XC": 90,
	"C": 100,
	//"CD": 400,
	"D":  500,
	"CM": 900,
	"M":  1000,
}

var combinedRomanDict = RomanDictionary{
	"IV": 4,
	"IX": 9,
	"XL": 40,
	"XC": 90,
	"CD": 400,
	"CM": 900,
}

func NewRomanDictionary() RomanDictionary {
	return RomanDictionary{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}
}
