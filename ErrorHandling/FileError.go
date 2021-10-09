package ErrorHandling

import "fmt"

func RecoverFromFileError() {
	fileError := recover()
	if fileError != nil {
		fmt.Printf("recovered from %v\n", fileError)
	}
}
