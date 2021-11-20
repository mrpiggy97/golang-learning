package InOut

import (
	"bufio"
	"fmt"
	"os"
)

func BufferReader() {
	//these program will take input from the user and write
	//to the console only when delimeter is entered by the user
	//in this case r
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	fmt.Println("<<< what do you have to say")
	fmt.Print("<<< ")
	text, err := reader.ReadString('r')
	if err != nil {
		panic(err)
	}
	fmt.Println(">>> you are right!!!!")
	fmt.Println(">>> " + text)
}
