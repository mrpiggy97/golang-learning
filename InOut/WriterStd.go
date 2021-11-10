package InOut

import (
	"fmt"
	"os"
)

func WriterStd() {
	//this program will write to the console
	var message []byte = []byte("this is using the standard library\n")
	bytesWritten, err := os.Stdout.Write(message)
	if err != nil {
		panic(err)
	}

	fmt.Printf("the number of bytes written is %v\n", bytesWritten)
}
