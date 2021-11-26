package InOut

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mrpiggy97/golang-learning/ErrorHandling"
)

// ReadFile will read and print the contents of given file.
func ReadFile(filePath string) {
	defer ErrorHandling.RecoverFromFileError()
	file, readError := ioutil.ReadFile(filePath)
	if readError != nil {
		panic(readError)
	}

	fmt.Printf("%s file\n", file)
}

// OsOffSet will read content of file starting at position 0.
func OsOffSet(filePath string) {
	file, fileError := os.OpenFile(filePath, os.O_RDWR, 0644)
	if fileError != nil {
		panic(fileError)
	}
	var positions []int64 = []int64{4}
	for _, position := range positions {
		_, err := file.Seek(position, 1)
		if err != nil {
			panic(err)
		}
	}

	file.WriteString("X")
	file.Seek(0, 0)
	fileBytes, fileBytesError := ioutil.ReadFile(filePath)
	if fileBytesError != nil {
		panic(fileBytesError)
	}
	fmt.Printf("%s\n", fileBytes)
	defer ErrorHandling.RecoverFromFileError()
	defer file.Close()
}
