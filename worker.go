package main

import (
	"fmt"
	"time"
)

// NewWorker creates, and returns a new Worker object. Its only argument
// is a channel that the worker can add itself to whenever it is done its
// work.
func NewWorker(id int, workerQueue chan chan CSPRequest, logger Logger) Worker {
	worker := Worker{
		ID:             id,
		Work:           make(chan CSPRequest),
		WorkerQueue:    workerQueue,
		QuitChan:       make(chan bool),
		Logger:         logger,
		CurrentWork:    make([]CSPRequest, 0),
		TimeoutStarted: false,
		Timeout:        time.Duration(5) * time.Second,
	}

	return worker
}

type Worker struct {
	ID             int
	Work           chan CSPRequest
	WorkerQueue    chan chan CSPRequest
	QuitChan       chan bool
	Logger         Logger
	CurrentWork    []CSPRequest
	TimeoutStarted bool
	Timeout        time.Duration
}

// This function "starts" the worker by starting a goroutine, that is
// an infinite "for-select" loop.
func (w *Worker) Start() {
	go func() {

		for {
			// Add ourselves into the worker queue.
			w.WorkerQueue <- w.Work

			select {
			case work := <-w.Work:
				// Receive a work request.
				fmt.Printf("Worker%d: Received work request.\n", w.ID)
				w.CurrentWork = append(w.CurrentWork, work)
				if w.TimeoutStarted == false {
					time.AfterFunc(w.Timeout, w.Flush)
					w.TimeoutStarted = true
				}

			case <-w.QuitChan:
				// We have been asked to stop.
				fmt.Printf("Worker%d stopping.\n", w.ID)
				return
			}
		}
	}()
}

// Stop tells the worker to stop listening for work requests.
// Note that the worker will only stop *after* it has finished its work.
func (w *Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

func (w *Worker) Flush() {
	fmt.Println("Flush")
	for _, work := range w.CurrentWork {
		w.Logger.Log(work)
	}
	w.TimeoutStarted = false
}
