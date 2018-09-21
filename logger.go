package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Logger interface {
	Log(data CSPRequest)
}

type StdoutLogger struct{}

func (l *StdoutLogger) Log(data CSPRequest) {
	b, _ := json.MarshalIndent(data, "", "    ")
	fmt.Println(string(b))
}

type CombinedLogger struct {
	Loggers []Logger
}

func (l *CombinedLogger) Log(data CSPRequest) {
	for _, logger := range l.Loggers {
		logger.Log(data)
	}
}

type HTTPLogger struct {
	Url string
}

func (l *HTTPLogger) Log(data CSPRequest) {
	jsn, err := json.Marshal(data.Report)
	if err != nil {
		panic(err)
	}

	request, err2 := http.NewRequest("POST", l.Url, bytes.NewBuffer(jsn))
	if err2 != nil {
		panic(err2)
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Response Status:", response.Status)
}
