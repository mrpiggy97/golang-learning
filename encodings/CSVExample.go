package encodings

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func CSVExample() {
	in := `user_id,score,password,"gopher",1000,admin,"BigJ",10,"1234","GGBoom",1111`

	var reader *strings.Reader = strings.NewReader(in)
	var csvReader *csv.Reader = csv.NewReader(reader)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		fmt.Println(record)
	}
}
