package astiworker

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/asticode/go-astilog"
)

// Worker represents an object capable of blocking, handling signals and stopping
type Worker struct {
	channelQuit chan bool
}

// NewWorker builds a new worker
func NewWorker() *Worker {
	astilog.Info("Starting Worker...")
	return &Worker{channelQuit: make(chan bool)}
}

// Close closes the worker
func (w Worker) Close() {
	astilog.Info("Closing Worker...")
}

// HandleSignals handles signals
func (w Worker) HandleSignals() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGABRT, syscall.SIGKILL, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	go func() {
		for s := range ch {
			astilog.Infof("Received signal %s", s)
			w.Stop()
			return
		}
	}()
}

// Stop stops the Worker
func (w *Worker) Stop() {
	astilog.Info("Stopping Worker...")
	if w.channelQuit != nil {
		close(w.channelQuit)
		w.channelQuit = nil
	}
}

// Wait is a blocking pattern
func (w *Worker) Wait() {
	astilog.Info("Worker is now waiting...")
	if w.channelQuit == nil {
		return
	}
	<-w.channelQuit
}
