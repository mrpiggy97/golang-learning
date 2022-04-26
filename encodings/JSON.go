package encodings

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type User struct {
	Username string `json:"username,omitempty"`
	Score    int    `json:"score,omitempty"`
	password string
}

func AsJSON(userInstance User) ([]byte, error) {
	jsonUser, marshalError := json.Marshal(userInstance)
	if marshalError != nil {
		return nil, nil
	}
	fmt.Println(string(jsonUser))
	return jsonUser, nil
}

func AsOriginal() {
	var newUser User = User{Username: "fabiaj", Score: 111, password: "mierda"}
	fmt.Println(newUser)
	jsonUser, marshalError := AsJSON(newUser)
	if marshalError != nil {
		panic(marshalError)
	}
	var normalUser *User = new(User)
	unmarshalError := json.Unmarshal(jsonUser, normalUser)
	if unmarshalError != nil {
		panic(unmarshalError)
	}
	fmt.Println(*normalUser)
}

func PrintJSON() {
	var userA User = User{Username: "fabian", Score: 222, password: "ps5"}
	var userB User = User{Username: "chris", Score: 999, password: "me lo confirmo"}
	var userC User = User{Username: "maddie", Score: 11111, password: "gente de zona"}
	var userSlice []User = []User{userA, userB, userC}
	jsonUserSlice, marshalError := json.Marshal(userSlice)
	if marshalError != nil {
		panic(marshalError)
	}
	var indent *bytes.Buffer = new(bytes.Buffer)
	indentErr := json.Indent(indent, jsonUserSlice, "", " ")
	if indentErr != nil {
		panic(indentErr)
	}

	fmt.Println(string(jsonUserSlice))
	fmt.Println(indent.String())
}
