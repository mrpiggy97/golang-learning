package InOut

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mrpiggy97/golang-learning/ErrorHandling"
)

func WriteToFile(sliceInstance []string, filePath string) {
	//ioutil.WriteFile overwrites all content
	defer ErrorHandling.RecoverFromFileError()
	for _, name := range sliceInstance {
		var writeError error = ioutil.WriteFile(filePath, []byte(name), 0644)
		if writeError != nil {
			panic(writeError)
		}
	}
}

func OsWriter(filePath string, message []string) {
	file, fileError := os.Create(filePath)
	if fileError != nil {
		panic(fileError)
	}
	for _, currentString := range message {
		file.WriteString(fmt.Sprintf("%v\n", currentString))
	}

	defer file.Close()
	defer ErrorHandling.RecoverFromFileError()
}
