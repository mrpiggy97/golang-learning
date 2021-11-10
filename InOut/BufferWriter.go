package InOut

import (
	"bufio"
	"os"
	"time"
)

func BufferWriter() {
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	var message string = "rule the world with go\n"
	for _, runeLetter := range message {
		time.Sleep(time.Millisecond * 100)
		writer.WriteRune(runeLetter)
		writer.Flush()
	}
}
