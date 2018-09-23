package main

import (
	"log"
)

var WorkerQueue chan chan CSPRequest

func StartDispatcher(nworkers int, logger Logger) {
	WorkerQueue = make(chan chan CSPRequest, nworkers)

	for i := 0; i < nworkers; i++ {
		log.Printf("Starting worker #%d.", i+1)
		worker := NewWorker(i+1, WorkerQueue, logger)
		worker.Start()
	}

	go func() {
		for {
			select {
			case work := <-WorkQueue:
				log.Print("Received work request.")
				go func() {
					worker := <-WorkerQueue

					log.Print("Dispatching work request.")
					worker <- work
				}()
			}
		}
	}()
}
