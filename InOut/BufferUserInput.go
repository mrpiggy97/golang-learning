package InOut

import (
	"bufio"
	"fmt"
	"os"
)

// BufferUserInput will just take input and print it.
func BufferUserInput() {
	fmt.Println(">>> please enter your input:")
	fmt.Print(">>> ")
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	//begin reading at specified delimiter
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Println(">>> your input was " + text)
}
