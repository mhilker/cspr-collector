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
	OutputHTTPHost    = flag.String("output-http-host", "http://localhost:9200/csp-violations/_doc", "http host to send the csp violations to")
	OutputESEnabled   = flag.Bool("output-es", false, "enable elasticsearch output")
	OutputESHost      = flag.String("output-es-host", "http://localhost:9200/", "elasticsearch host to send the csp violations to")
)

func main() {
	// Parse the command-line flags.
	flag.Parse()

	loggers := []Logger{}

	if *OutputStdout {
		log.Printf("Enable Stdout Logger")
		loggers = append(loggers, &StdoutLogger{})
	}
	if *OutputHTTPEnabled {
		log.Printf("Enable HTTP Logger")
		loggers = append(loggers, &HTTPLogger{Url: *OutputHTTPHost})
	}
	if *OutputESEnabled {
		log.Printf("Enable ES Logger")
		loggers = append(loggers, &ElasticsearchLogger{Url: *OutputESHost})
	}

	logger := &CombinedLogger{Loggers: loggers}

	// Start the dispatcher.
	log.Println("Starting the dispatcher")
	StartDispatcher(*NWorkers, logger)

	// Register our collector as an HTTP handler function.
	log.Println("Registering the collector")
	http.HandleFunc("/", Collector)

	// Start the HTTP server!
	log.Println("HTTP server listening on", *HTTPListenHost)
	if err := http.ListenAndServe(*HTTPListenHost, nil); err != nil {
		log.Println(err.Error())
	}
}
