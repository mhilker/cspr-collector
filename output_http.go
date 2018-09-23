package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type HTTPOutput struct {
	Url string
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
