package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var WorkQueue = make(chan CSPRequest, 100)

func Collector(response http.ResponseWriter, request *http.Request) {
	if request.RequestURI != "/" {
		response.WriteHeader(http.StatusNotFound)
		log.Printf("Path \"%s\" not found.\n", request.RequestURI)
		return
	}

	if request.Method != "POST" {
		response.Header().Set("Allow", "POST")
		response.WriteHeader(http.StatusMethodNotAllowed)
		log.Printf("Method \"%s\" not allowed.\n", request.Method)
		return
	}

	contentType := request.Header.Get("Content-Type")
	if contentType != "application/json" {
		response.WriteHeader(http.StatusUnsupportedMediaType)
		log.Printf("Unsupported Media Type \"%s\".\n", contentType)
		return
	}

	body, err1 := ioutil.ReadAll(request.Body)
	if err1 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		log.Println(err1.Error())
		return
	}

	data := NewCSPRequest()
	err2 := json.Unmarshal(body, &data)
	if err2 != nil {
		response.WriteHeader(http.StatusBadRequest)
		log.Println(err1.Error())
		return
	}

	WorkQueue <- data
	log.Println("CSPRequest queued.")

	response.WriteHeader(http.StatusCreated)
	return
}
