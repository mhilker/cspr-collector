package main

import (
	"log"
	"time"
)

func NewWorker(id int, workerQueue chan chan CSPRequest, output Output) Worker {
	worker := Worker{
		ID:             id,
		Work:           make(chan CSPRequest),
		WorkerQueue:    workerQueue,
		QuitChan:       make(chan bool),
		Output:         output,
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
	Output         Output
	CurrentWork    []CSPRequest
	TimeoutStarted bool
	Timeout        time.Duration
}

func (w *Worker) Start() {
	go func() {

		for {
			w.WorkerQueue <- w.Work

			select {
			case work := <-w.Work:
				log.Printf("Worker #%d: Received work request.", w.ID)
				w.CurrentWork = append(w.CurrentWork, work)
				if w.TimeoutStarted == false {
					time.AfterFunc(w.Timeout, w.Flush)
					w.TimeoutStarted = true
				}

			case <-w.QuitChan:
				log.Printf("Worker #%d stopping.", w.ID)
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

func (w *Worker) Flush() {
	log.Printf("Flush %d entries.", len(w.CurrentWork))
	w.Output.Write(w.CurrentWork)
	w.TimeoutStarted = false
}
