package InOut

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func Write() {
	//this program will write to the console
	var message []byte = []byte("this is using the standard library\n")
	bytesWritten, err := os.Stdout.Write(message)
	if err != nil {
		panic(err)
	}

	fmt.Printf("the number of bytes written is %v\n", bytesWritten)
}

func ReaderStd() {
	//this function will set a set limit of bytes for what
	//can be written to the console or terminal
	//if this limit is exeeded then the console will just take
	//the number of characters allowed and dismiss the rest
	//provoking a loss of information
	var target []byte = make([]byte, 3)
	bytesRead, err := os.Stdin.Read(target)
	if err != nil {
		panic(err)
	}
	var message string = string(target[0:bytesRead])
	var printedMessage string = fmt.Sprintf("%v %v", bytesRead, strings.ToUpper(message))
	for _, runeLetter := range printedMessage {
		time.Sleep(time.Millisecond * 300)
		fmt.Print(runeLetter)
	}
}
