package csprcollector

import (
	"encoding/json"
	"log"
)

type StdoutOutput struct{}

func (o *StdoutOutput) Write(data []CSPRequest) {
	for _, d := range data {
		b, _ := json.MarshalIndent(d, "", "    ")
		log.Print(string(b))
	}
}
