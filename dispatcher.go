package csprcollector

import (
	"log"
)

func NewDispatcher(nworkers int, output Output, workQueue <-chan CSPRequest) *Dispatcher {
	return &Dispatcher{
		WorkerQueue:     make(chan chan CSPRequest, nworkers),
		WorkQueue:       workQueue,
		NumberOfWorkers: nworkers,
		Output:          output,
	}
}

type Dispatcher struct {
	WorkerQueue     chan chan CSPRequest
	WorkQueue       <-chan CSPRequest
	NumberOfWorkers int
	Output          Output
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.NumberOfWorkers; i++ {
		log.Printf("Starting worker #%d.", i+1)
		worker := NewWorker(i+1, d.WorkerQueue, d.Output)
		worker.Start()
	}

	go d.start()
}

func (d *Dispatcher) start() {
	for {
		select {
		case work := <-d.WorkQueue:
			log.Print("Received work request.")
			go func() {
				worker := <-d.WorkerQueue
				log.Print("Dispatching work request.")
				worker <- work
			}()
		}
	}
}
