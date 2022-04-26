package encodings

import (
	"encoding/csv"
	"os"
)

func CSVWrite() {
	var output [][]string = [][]string{
		{"user_id", "score", "password"},
		{"gopher", "1000", "admin,"},
		{"Bigj", "10", "1234"},
		{"GGBoom", "", "1111"},
	}

	var writer *csv.Writer = csv.NewWriter(os.Stdout)
	for _, record := range output {
		var err error = writer.Write(record)
		if err != nil {
			panic(err)
		}
	}

	writer.Flush()
}
