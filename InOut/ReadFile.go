package InOut

import (
	"fmt"
	"io/ioutil"

	"github.com/mrpiggy97/golang-learning/ErrorHandling"
)

func ReadFile(filePath string) {
	defer ErrorHandling.RecoverFromFileError()
	file, readError := ioutil.ReadFile(filePath)
	if readError != nil {
		panic(readError)
	}

	fmt.Printf("%s file\n", file)
}
