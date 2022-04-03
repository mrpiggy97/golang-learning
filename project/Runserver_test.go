package project_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/mrpiggy97/golang-learning/project"
)

func TestServer(testCase *testing.T) {
	go project.Runserver()
	var requestJob project.Job = project.Job{
		Name:   "chris",
		Delay:  time.Second * 1,
		Number: 11212,
	}
	jsonData, _ := json.Marshal(requestJob)
	var buferrer *bytes.Buffer = bytes.NewBuffer(jsonData)
	request, _ := http.NewRequest(
		"POST",
		"http://localhost:8081/fib",
		buferrer,
	)
	var client *http.Client = &http.Client{}
	res, resErr := client.Do(request)
	if resErr != nil {
		testCase.Errorf("expected response error to be nil")
	}
	fmt.Println(res.Status)
}
