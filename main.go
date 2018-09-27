package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	NWorkers          = flag.Int("n", 4, "the number of workers to start")
	HTTPListenHost    = flag.String("host", "127.0.0.1:8080", "address to listen for http requests on")
	OutputStdout      = flag.Bool("output-stdout", true, "enable stdout output")
	OutputHTTPEnabled = flag.Bool("output-http", false, "enable http output")
	OutputHTTPHost    = flag.String("output-http-host", "http://localhost:80/", "http host to send the csp violations to")
	OutputESEnabled   = flag.Bool("output-es", false, "enable elasticsearch output")
	OutputESHost      = flag.String("output-es-host", "http://localhost:9200/", "elasticsearch host to send the csp violations to")
	OutputESIndex     = flag.String("output-es-index", "csp-violations", "elasticsearch index to save the csp violations in")
)

func main() {
	flag.Parse()

	var outputs []Output

	if *OutputStdout {
		log.Printf("Enable Stdout Output.")
		outputs = append(outputs, &StdoutOutput{})
	}
	if *OutputHTTPEnabled {
		log.Printf("Enable HTTP Output.")
		outputs = append(outputs, &HTTPOutput{Url: *OutputHTTPHost})
	}
	if *OutputESEnabled {
		log.Printf("Enable ES Output.")
		outputs = append(outputs, &ElasticsearchOutput{
			Url:   *OutputESHost,
			Index: *OutputESIndex,
		})
	}

	output := &CombinedOutput{Outputs: outputs}

	StartDispatcher(*NWorkers, output)
	http.HandleFunc("/", Collector)

	log.Printf("HTTP server listening on %s.", *HTTPListenHost)
	err := http.ListenAndServe(*HTTPListenHost, nil)
	if err != nil {
		log.Print(err.Error())
	}
}
