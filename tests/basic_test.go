package tests

import "testing"

//TestBasic is just a simple test.
func TestBasic(testCase *testing.T) {
	var total int = 2 + 2
	if total != 4 {
		testCase.Error("expected 4 and got ", total)
	}
}
