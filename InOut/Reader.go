package InOut

import (
	"errors"
	"fmt"
	"io"
)

// Reader is just an example of a reader.
type Reader struct {
	Data string
	// From will tell Read function where to start reading.
	From int
}

// Read will swap the values of byteSlice with the byte value of members
// in readerInstance.Data
// for that will set a number of items to read, this number will initially come from
// substracting the length of readerInstance.Data and readerInstance.From
// however if the length of byteSlice is lower than this number we will just set this
// number to the length of the byteSlice.
func (readerInstance *Reader) Read(byteSlice []byte) (int, error) {

	if byteSlice == nil {
		return -1, errors.New("nil target slice")
	}
	if len(readerInstance.Data) <= 0 || readerInstance.From == len(readerInstance.Data) {
		return 0, io.EOF
	}

	//define starting point for reader
	var numberOfItemsToRead int = len(readerInstance.Data) - readerInstance.From
	fmt.Printf("%v this is the number of items to read\n", numberOfItemsToRead)

	//if byteSlice length is less than the number of items to read
	//then set the number of items to read to the length of byteSlice
	if len(byteSlice) < numberOfItemsToRead {
		numberOfItemsToRead = len(byteSlice)
	}

	for i := 0; i < numberOfItemsToRead; i++ {
		var newValueForIndex byte = readerInstance.Data[readerInstance.From]
		byteSlice[i] = newValueForIndex
		readerInstance.From++
	}

	if readerInstance.From == len(readerInstance.Data) {
		return numberOfItemsToRead, io.EOF
	}

	return numberOfItemsToRead, nil
}
