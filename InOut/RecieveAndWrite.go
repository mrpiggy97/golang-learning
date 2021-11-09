package InOut

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/mrpiggy97/golang-learning/ErrorHandling"
)

type CreatedFunc func()

func InsertAtPosition(index int, slice []string, newMember string) []string {
	if index >= len(slice) {
		var newError error = errors.New("index cannot equal or be greater than length of slice")
		panic(newError)
	}

	var copySlice []string = []string{}
	copySlice = append(copySlice, slice[0:index]...)
	copySlice = append(copySlice, newMember)
	copySlice = append(copySlice, slice[index:]...)

	return copySlice
}

func RecieveAndWrite(channel <-chan map[string]string, filePath string, waiter *sync.WaitGroup) {
	waiter.Add(1)
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	var Functions []string = []string{"{", "}"}
	file.WriteString(`
	package InOut
	import "fmt"
	`)
	for {
		funcInfo, channelIsAvailable := <-channel
		if !channelIsAvailable {
			break
		}
		Functions = InsertAtPosition(1, Functions, funcInfo["funcName"])
		file.WriteString(fmt.Sprintf("%v\n", funcInfo["func"]))
	}
	fmt.Printf("%v\n", Functions)
	file.WriteString(fmt.Sprintf("var Fs []CreatedFunc = []CreatedFunc %v", strings.Join(Functions, "")))
	defer waiter.Done()
	defer ErrorHandling.RecoverFromFileError()
	defer file.Close()
}
