package InOut

import (
	"bufio"
	"os"
	"time"
)

// BufferWriter will allow user to write input and then
// it will print out a different thing just because.
func BufferWriter() {
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	var message string = "rule the world with go\n"
	for _, runeLetter := range message {
		time.Sleep(time.Millisecond * 100)
		writer.WriteRune(runeLetter)
		writer.Flush()
	}
}
