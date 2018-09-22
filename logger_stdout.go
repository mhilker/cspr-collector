package main

import (
	"encoding/json"
	"log"
)

type StdoutLogger struct{}

func (l *StdoutLogger) Log(data []CSPRequest) {
	for _, d := range data {
		b, _ := json.MarshalIndent(d, "", "    ")
		log.Println(string(b))
	}
}
