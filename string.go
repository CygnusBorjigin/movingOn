package MovingOn

import (
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

func NumericInequality(target string) (bool, bool, bool) {
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
			if firstChar == '>' && secondChar == '=' {
				return true, true, true
			} else if firstChar == '<' && secondChar == '=' {
				return true, false, true
			} else if firstChar == '>' && secondChar != '=' {
				return true, true, false
			} else if firstChar == '<' && secondChar != '=' {
				return true, false, false
			}
		}
	} else {
		if NumericString(restChar1) {
			if firstChar == '>' {
				return true, true, false
			} else if firstChar == '<' {
				return true, false, false
			}
		}
	}

	return false, false, false
}
