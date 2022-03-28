package stringManipulation

import (
	"strings"
)

// StringToUpperCase will turn every character of instance
// to Upper only if that character is lower
func StringToUpperCase(instance string) string {

	var convertedString string = ""
	for _, character := range instance {
		var testString string = strings.ToUpper(string(character))
		if testString == string(character) {
			convertedString = convertedString + string(character)
		} else {
			var convertedCharacter string = strings.ToUpper(string(character))
			convertedString = convertedString + convertedCharacter
		}
	}

	return convertedString
}
