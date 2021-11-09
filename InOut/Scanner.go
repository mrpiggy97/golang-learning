package InOut

import (
	"bufio"
	"fmt"
	"os"
)

func Scanner() {
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	fmt.Println(">>> what is the input?")
	var counter int = 0
	for scanner.Scan() {
		var text string = scanner.Text()
		counter = counter + len(text)
		if counter > 15 {
			break
		}
	}

	fmt.Println("that's enough")
}
