package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	NWorkers   = flag.Int("n", 4, "The number of workers to start")
	HTTPAddr   = flag.String("http", "127.0.0.1:8080", "Address to listen for HTTP requests on")
	HTTPTarget = flag.String("target", "http://localhost:9200/csp-violations/_doc", "Host to send the csp violations to")
)

func main() {
	// Parse the command-line flags.
	flag.Parse()

	stdoutLogger := &StdoutLogger{}
	httpLogger := &HTTPLogger{Url: *HTTPTarget}
	logger := &CombinedLogger{Loggers: []Logger{stdoutLogger, httpLogger}}

	// Start the dispatcher.
	fmt.Println("Starting the dispatcher")
	StartDispatcher(*NWorkers, logger)

	// Register our collector as an HTTP handler function.
	fmt.Println("Registering the collector")
	http.HandleFunc("/", Collector)

	// Start the HTTP server!
	fmt.Println("HTTP server listening on", *HTTPAddr)
	if err := http.ListenAndServe(*HTTPAddr, nil); err != nil {
		fmt.Println(err.Error())
	}
}
