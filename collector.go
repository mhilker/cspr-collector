package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var WorkQueue = make(chan CSPRequest, 100)

func Collector(response http.ResponseWriter, request *http.Request) {
	if request.RequestURI != "/" {
		response.WriteHeader(http.StatusNotFound)
		fmt.Printf("Path \"%s\" not found.\n", request.RequestURI)
		return
	}

	if request.Method != "POST" {
		response.Header().Set("Allow", "POST")
		response.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Printf("Method \"%s\" not allowed.\n", request.Method)
		return
	}

	contentType := request.Header.Get("Content-Type")
	if contentType != "application/json" {
		response.WriteHeader(http.StatusUnsupportedMediaType)
		fmt.Printf("Unsupported Media Type \"%s\".\n", contentType)
		return
	}

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	data := NewCSPRequest()
	err2 := json.Unmarshal(body, &data)
	if err2 != nil {
		response.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	WorkQueue <- data
	fmt.Println("CSPRequest queued.")

	response.WriteHeader(http.StatusCreated)
	return
}
