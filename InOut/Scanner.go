package InOut

import (
	"bufio"
	"fmt"
	"os"
)

// Scanner take input from user and prints it out to the console.
func Scanner() {
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	fmt.Println(">>> what is the input?")
	var counter int = 0
	fmt.Print("<<< ")
	for scanner.Scan() {
		var text string = scanner.Text()
		counter = counter + len(text)
		if counter > 15 {
			break
		} else {
			fmt.Print("<<< ")
		}
	}

	fmt.Println("that's enough")
}
