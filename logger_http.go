package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type HTTPLogger struct {
	Url string
}

func (l *HTTPLogger) Log(data []CSPRequest) {
	for _, d := range data {
		jsn, err := json.Marshal(d.Report)
		if err != nil {
			log.Println(err.Error())
			return
		}

		request, err2 := http.NewRequest("POST", l.Url, bytes.NewBuffer(jsn))
		if err2 != nil {
			log.Println(err2.Error())
			return
		}
		request.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			log.Println(err.Error())
			return
		}
		defer response.Body.Close()

		log.Println("Response Status:", response.Status)
	}
}
