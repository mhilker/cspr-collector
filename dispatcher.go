package main

import (
	"log"
)

var WorkerQueue chan chan CSPRequest

func StartDispatcher(nworkers int, logger Logger) {
	// First, initialize the channel we are going to but the workers' work channels into.
	WorkerQueue = make(chan chan CSPRequest, nworkers)

	// Now, create all of our workers.
	for i := 0; i < nworkers; i++ {
		log.Println("Starting worker", i+1)
		worker := NewWorker(i+1, WorkerQueue, logger)
		worker.Start()
	}

	go func() {
		for {
			select {
			case work := <-WorkQueue:
				log.Println("Received work requeust")
				go func() {
					worker := <-WorkerQueue

					log.Println("Dispatching work request")
					worker <- work
				}()
			}
		}
	}()
}
