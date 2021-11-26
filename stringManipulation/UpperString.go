package stringManipulation

import (
	"fmt"
	"strings"
)

// StringToUpperCase will turn every character of instance
// to Upper only if that character is lower
func StringToUpperCase(instance string) {

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

	fmt.Printf("%v string\n", convertedString)
}
