package InOut

import (
	"fmt"
	"os"
)

// WriterStd will write to the console.
func WriterStd() {
	var message []byte = []byte("this is using the standard library\n")
	bytesWritten, err := os.Stdout.Write(message)
	if err != nil {
		panic(err)
	}

	fmt.Printf("the number of bytes written is %v\n", bytesWritten)
}
