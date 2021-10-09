package InOut

import (
	"io/ioutil"

	"github.com/mrpiggy97/golang-learning/ErrorHandling"
)

func WriteToFile(sliceInstance []string, filePath string) {
	defer ErrorHandling.RecoverFromFileError()
	for _, name := range sliceInstance {
		var writeError error = ioutil.WriteFile(filePath, []byte(name), 0644)
		if writeError != nil {
			panic(writeError)
		}
	}
}
