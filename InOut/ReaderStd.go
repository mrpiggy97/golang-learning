package InOut

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func ReaderStd() {
	//this function will set a set limit of bytes for what
	//can be written to the console or terminal
	//if this limit is exeeded then the console will just take
	//the number of characters allowed and dismiss the rest
	//provoking a loss of information
	var target []byte = make([]byte, 3)
	fmt.Println("<<< enter value")
	bytesRead, err := os.Stdin.Read(target)
	if err != nil {
		panic(err)
	}
	var message string = string(target[0:bytesRead])
	var printedMessage string = fmt.Sprintf("%v %v", bytesRead, strings.ToUpper(message))
	for _, runeLetter := range printedMessage {
		time.Sleep(time.Millisecond * 300)
		var runeString string = fmt.Sprintf("%v ", runeLetter)
		fmt.Println(runeString)
	}
}
