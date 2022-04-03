package project

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

func RequestHandler(writer http.ResponseWriter, request *http.Request, jobQueue chan Job) {
	if request.Method != "POST" {
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}

	jsonData, _ := io.ReadAll(request.Body)
	var jobData *Job = new(Job)
	var unMarshalingError error = json.Unmarshal(jsonData, jobData)
	if unMarshalingError != nil {
		log.Error().Msg(unMarshalingError.Error())
	}
	jobQueue <- *jobData
}

func Runserver() {
	const (
		maxWorkers   = 4
		maxQueueSize = 20
		port         = "0.0.0.0:8081"
	)
	var jobQueue chan Job = make(chan Job, maxQueueSize)
	var dispatcher *Dispatcher = NewDispatcher(jobQueue, maxWorkers)
	dispatcher.Run()
	http.HandleFunc("/fib", func(writer http.ResponseWriter, req *http.Request) {
		RequestHandler(writer, req, jobQueue)
	})
	http.ListenAndServe(port, nil)
}
