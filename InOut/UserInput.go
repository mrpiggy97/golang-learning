package InOut

import (
	"bufio"
	"fmt"
	"os"
)

func UserInput() {
	fmt.Println("please enter message")
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var text string = scanner.Text()
	fmt.Println(text)
}
