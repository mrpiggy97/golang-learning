package stringManipulation

import "testing"

func MakeStringUpperCase(testCase *testing.T) {
	var stringToConvert string = "MyString"
	stringToConvert = StringToUpperCase(stringToConvert)
	if stringToConvert != "MYSTRING" {
		testCase.Error("expected stringtoConvert to be MYSTRING")
	}
}

func TestUpperString(testCase *testing.T) {
	testCase.Run("action=test-string-to-upper-case", MakeStringUpperCase)
}
