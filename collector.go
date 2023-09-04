package csprcollector

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func NewCollector(workQueue chan<- CSPRequest) *Collector {
	return &Collector{
		WorkQueue: workQueue,
	}
}

type Collector struct {
	WorkQueue chan<- CSPRequest
}

func (c *Collector) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/health" {
		c.handleHealth(w)
		return
	}

	if r.RequestURI != "/" {
		message := fmt.Sprintf("Path \"%s\" not found.", r.RequestURI)
		c.response(w, http.StatusNotFound, message)
		return
	}

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		message := fmt.Sprintf("Method \"%s\" not allowed.", r.Method)
		c.response(w, http.StatusMethodNotAllowed, message)
		return
	}

	contentType := r.Header.Get("Content-Type")
	if contentType != "application/csp-report" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		message := fmt.Sprintf("Unsupported Media Type \"%s\".", contentType)
		c.response(w, http.StatusUnsupportedMediaType, message)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		message := err.Error()
		c.response(w, http.StatusInternalServerError, message)
		return
	}

	data := NewCSPRequest()
	err = json.Unmarshal(body, &data)
	if err != nil {
		message := err.Error()
		c.response(w, http.StatusBadRequest, message)
		return
	}

	c.WorkQueue <- data

	c.response(w, http.StatusCreated, "")
}

func (c *Collector) response(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)

	if message == "" {
		return
	}

	log.Print(message)

	m := map[string]string{
		"message": message,
	}

	j, _ := json.MarshalIndent(m, "", "    ")

	w.Header().Set("content-type", "application/json")
	_, err := w.Write(j)
	if err != nil {
		log.Fatal(err)
	}
}

func (c *Collector) handleHealth(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
