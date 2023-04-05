package MovingOn

import (
	"strconv"
	"strings"
)

func NumericString(target string) bool {
	for _, value := range strings.Split(target, "") {
		numbers := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
		if !StringSliceContainsString(numbers, value) {
			return false
		}
	}
	return true
}

func AlphabetString(target string) bool {
	for _, value := range strings.Split(target, "") {
		letters := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
		if !StringSliceContainsString(letters, value) {
			return false
		}
	}
	return true
}

func StringToInt(target string) (int64, *string) {
	if target[0] == '-' {
		parsedRes, parseErr := strconv.ParseInt(target[1:], 10, 64)
		if parseErr != nil {
			errorMessage := "Parse failed"
			return 0, &errorMessage
		}
		return -1 * parsedRes, nil
	} else {
		parsedRes, parseErr := strconv.ParseInt(target, 10, 64)
		if parseErr != nil {
			errorMessage := "Parse failed"
			return 0, &errorMessage
		}
		return parsedRes, nil
	}
}

func NumericInequality(target string, isNumericComparison bool, isGreaterThan bool, isEqualTo bool, isInequality bool, numericValue int64) (bool, bool, bool, bool, int64) {
	// Meaning of the return value
	//  1. Whether the string represents a numerical comparison
	//  2. Whether the inequality is a greater than relationship
	//  3. Whether the inequality is an equal to relationship
	//  4. Whether the comparison is a inequality comparison
	if NumericString(target) {
		parsedValue, _ := StringToInt(target)
		return isNumericComparison, isGreaterThan, isEqualTo, isInequality, parsedValue
	}
	switch target[0] {
	case '>':
		return NumericInequality(target[1:], true, true, isEqualTo, true, -1)
	case '<':
		return NumericInequality(target[1:], true, false, isEqualTo, true, -1)
	case '=':
		return NumericInequality(target[1:], true, isGreaterThan, true, isInequality, -1)
	default:
		return false, false, false, false, -1
	}
}
