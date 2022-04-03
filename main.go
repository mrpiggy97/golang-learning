package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/mrpiggy97/golang-learning/project"
	"github.com/rs/zerolog/log"
)

// Main will apply all concepts learned.

func RequestHandler(writer http.ResponseWriter, request *http.Request, jobQueue chan project.Job) {
	if request.Method != "POST" {
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}

	jsonData, _ := io.ReadAll(request.Body)
	var jobData *project.Job = new(project.Job)
	var unMarshalingError error = json.Unmarshal(jsonData, jobData)
	if unMarshalingError != nil {
		log.Error().Msg(unMarshalingError.Error())
	}
	jobQueue <- *jobData
}

func main() {
	const (
		maxWorkers   = 4
		maxQueueSize = 20
		port         = "0.0.0.0:8081"
	)
	var jobQueue chan project.Job = make(chan project.Job, maxQueueSize)
	var dispatcher *project.Dispatcher = project.NewDispatcher(jobQueue, maxWorkers)
	dispatcher.Run()
	http.HandleFunc("/fib", func(writer http.ResponseWriter, req *http.Request) {
		RequestHandler(writer, req, jobQueue)
	})
	http.ListenAndServe(port, nil)
}
