package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"

	cspr "github.com/mhilker/cspr-collector"
)

var (
	NumberOfWorkers   = flag.Int("n", 4, "the number of workers to start")
	HTTPListenHost    = flag.String("host", "127.0.0.1:8080", "address to listen for http requests on")
	OutputStdout      = flag.Bool("output-stdout", false, "enable stdout output")
	OutputHTTPEnabled = flag.Bool("output-http", false, "enable http output")
	OutputHTTPHost    = flag.String("output-http-host", "http://localhost:80/", "http host to send the csp violations to")
	OutputHTTPHeaders = flag.String("output-http-headers", "", "additional headers for HTTP output")
	OutputESEnabled   = flag.Bool("output-es", false, "enable elasticsearch output")
	OutputESHost      = flag.String("output-es-host", "http://localhost:9200/", "elasticsearch host to send the csp violations to")
	OutputESIndex     = flag.String("output-es-index", "cspr-violations", "elasticsearch index to save the csp violations in")
	OutputEsCertFile  = flag.String("output-es-cert-file", "", "cert file for elasticsearch")
	OutputEsKeyFile   = flag.String("output-es-key-file", "", "key file for elasticsearch")
	OutputEsCaFile    = flag.String("output-es-ca-file", "", "ca file for elasticsearch")
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	flag.Parse()

	workQueue := make(chan cspr.CSPRequest, 100)

	output := NewOutput()
	dispatcher := cspr.NewDispatcher(*NumberOfWorkers, output, workQueue)
	dispatcher.Run()

	collector := cspr.NewCollector(workQueue)
	server := &http.Server{Addr: *HTTPListenHost, Handler: collector}

	go func() {
		log.Printf("HTTP server listening on %s.", *HTTPListenHost)
		if err := server.ListenAndServe(); err != nil {
			log.Print(err.Error())
		}
	}()

	<-stop

	log.Print("Shutting down the server.")
	err := server.Shutdown(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("Server gracefully stopped.")
}

func NewOutput() *cspr.CombinedOutput {
	var outputs []cspr.Output

	if *OutputStdout {
		log.Printf("Enable Stdout Output.")
		outputs = append(outputs, &cspr.StdoutOutput{})
	}

	if *OutputHTTPEnabled {
		log.Printf("Enable HTTP Output.")
		headers := cspr.ParseHeaders(*OutputHTTPHeaders)

		outputs = append(outputs, &cspr.HTTPOutput{
			Url:     *OutputHTTPHost,
			Headers: headers,
		})
	}

	if *OutputESEnabled {
		log.Printf("Enable ES Output.")
		outputs = append(outputs, &cspr.ElasticsearchOutput{
			Url:    *OutputESHost,
			Index:  *OutputESIndex,
			Client: cspr.NewHttpClient(*OutputEsCertFile, *OutputEsKeyFile, *OutputEsCaFile),
		})
	}

	return &cspr.CombinedOutput{Outputs: outputs}
}
