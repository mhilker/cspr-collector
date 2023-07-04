package csprcollector

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type HTTPOutput struct {
	Url     string
	Headers map[string]string
}

func (o *HTTPOutput) Write(data []CSPRequest) {
	for _, d := range data {
		jsn, err := json.Marshal(d.Report)
		if err != nil {
			log.Print(err.Error())
			return
		}

		request, err := http.NewRequest("POST", o.Url, bytes.NewBuffer(jsn))
		if err != nil {
			log.Print(err.Error())
			return
		}

		for key, value := range o.Headers {
			request.Header.Set(key, value)
		}
		request.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			log.Print(err.Error())
			return
		}

		defer response.Body.Close()
		log.Print("Response Status:", response.Status)
	}
}

func ParseHeaders(headerString string) map[string]string {
	headers := make(map[string]string)
	pairs := strings.Split(headerString, ",")

	for _, pair := range pairs {
		split := strings.SplitN(pair, ":", 2)
		if len(split) == 2 {
			key := strings.TrimSpace(split[0])
			value := strings.TrimSpace(split[1])
			headers[key] = value
		}
	}

	return headers
}
