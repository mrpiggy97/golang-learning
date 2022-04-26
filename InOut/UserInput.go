package InOut

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mrpiggy97/golang-learning/ErrorHandling"
)

// UserInput reads input from user and prints it out to console.s
func UserInput() {
	fmt.Println("<<< please enter message")
	fmt.Print("<<< ")
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var text string = scanner.Text()
	var filePath string = "/workspaces/golang-learning/InOut/user-input.txt"
	file, fileError := os.Create(filePath)
	if fileError != nil {
		panic(fileError)
	}
	_, writeError := file.WriteString(text)
	if writeError != nil {
		panic(writeError)
	}
	defer ErrorHandling.RecoverFromFileError()
}
