package InOut

import (
	"errors"
	"fmt"
	"io"
)

type Writer struct {
	Data string
	Size int
}

func (writerInstance *Writer) Write(target []byte) (int, error) {
	if len(target) == 0 {
		return 0, io.EOF
	}

	var currentLengthUsed int = writerInstance.Size

	var err error = nil
	if len(target) < currentLengthUsed {
		currentLengthUsed = len(target)
	} else {
		err = errors.New("target length is greater than writer.Size")
	}

	writerInstance.Data = writerInstance.Data + string(target[0:currentLengthUsed])
	fmt.Printf("writer data %v\n", writerInstance.Data)
	return currentLengthUsed, err
}
