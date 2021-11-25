package tests

func ChooseReturn(number int) string {
	if number > 100 {
		return "number is bigger than 100"
	} else {
		return "number is lower than 100"
	}
}
