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

func NumericInequality(target string) (bool, bool, bool, int64) {
	// Meaning of the return value
	//  1. Whether the string represents a numerical inequality
	//  2. Whether the inequality is a greater than relationship
	//  3. Whether the inequality is an equal to relationship

	firstChar := target[0]
	restChar1 := target[1:len(target)]
	if len(target) > 2 {
		secondChar := target[1]
		restChar2 := target[2:len(target)]
		if NumericString(restChar2) {
			parsedNumber, parseSuccess := StringToInt(restChar2)
			if parseSuccess != nil {
				return false, false, false, -1
			}
			if firstChar == '>' && secondChar == '=' {
				return true, true, true, parsedNumber
			} else if firstChar == '<' && secondChar == '=' {
				return true, false, true, parsedNumber
			} else if firstChar == '>' && secondChar != '=' {
				return true, true, false, parsedNumber
			} else if firstChar == '<' && secondChar != '=' {
				return true, false, false, parsedNumber
			}
		}
	} else {
		if NumericString(restChar1) {
			parsedNumber, parseError := StringToInt(restChar1)
			if parseError != nil {
				return false, false, false, -1
			}
			if firstChar == '>' {
				return true, true, false, parsedNumber
			} else if firstChar == '<' {
				return true, false, false, parsedNumber
			}
		}
	}

	return false, false, false, -1
}
