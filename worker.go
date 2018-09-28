package main

import (
	"log"
	"time"
)

func NewWorker(id int, workerQueue chan chan CSPRequest, output Output) Worker {
	return Worker{
		ID:          id,
		Work:        make(chan CSPRequest),
		WorkerQueue: workerQueue,
		Output:      output,
	}
}

type Worker struct {
	ID          int
	Work        chan CSPRequest
	WorkerQueue chan chan CSPRequest
	Output      Output
}

func (w *Worker) Start() {
	go func() {
		w.requeue()

		var buffer []CSPRequest
		ticker := time.NewTicker(time.Second * time.Duration(5))

		for {
			select {
			case work := <-w.Work:
				log.Printf("Worker #%d: Received work request.", w.ID)
				buffer = append(buffer, work)

				if len(buffer) >= 50 {
					log.Printf("Worker #%d: Buffer Flush.", w.ID)
					w.Flush(buffer)
					buffer = nil
				}
				w.requeue()
			case <-ticker.C:
				if len(buffer) > 0 {
					log.Printf("Worker #%d: Ticked Flush.", w.ID)
					w.Flush(buffer)
					buffer = nil
				}
			}
		}
	}()
}

func (w *Worker) Flush(requests []CSPRequest) {
	log.Printf("Flush %d entries.", len(requests))
	w.Output.Write(requests)
}

func (w *Worker) requeue() {
	w.WorkerQueue <- w.Work
}
