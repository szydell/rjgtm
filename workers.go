package main

import (
	"sync"

	"github.com/cenkalti/rpc2"
	rjerr "github.com/szydell/rjgtm/rjerr"
)

type allWorkers struct {
	sync.RWMutex
	list map[*rpc2.Client]bool
}

type glvn struct {
	key   string
	value string
}

//NewAllWorkers constructor to initialize map
func newAllWorkers() *allWorkers {
	var w allWorkers
	w.list = make(map[*rpc2.Client]bool)
	return &w
}

func (w *allWorkers) getState(worker *rpc2.Client) bool {
	w.RLock()
	defer w.RUnlock()
	return w.list[worker]
}

func (w *allWorkers) setState(worker *rpc2.Client, state bool) {
	w.Lock()
	defer w.Unlock()
	w.list[worker] = state
}

func (w *allWorkers) subscribeWorker(worker *rpc2.Client) {
	w.Lock()
	defer w.Unlock()
	w.list[worker] = true
}

func (w *allWorkers) unSubscribeWorker(worker *rpc2.Client) {
	w.Lock()
	defer w.Unlock()
	delete(w.list, worker)
}

func (w *allWorkers) returnWorkers() map[*rpc2.Client]bool {
	w.RLock()
	defer w.RUnlock()
	return w.list
}

func (w *allWorkers) doWork(command string, data interface{}) (reply string, err error) {
	if len(w.list) < 1 {
		return "", rjerr.ErrNoAvailableWorkers
	}
	var worker *rpc2.Client
	var state bool
	w.Lock()
	for worker, state = range w.list {
		if state {
			break
		}
	}
	if worker == nil {
		w.Unlock()
		return "", rjerr.ErrAllWorkersBusy
	}
	w.list[worker] = false
	w.Unlock()
	err = worker.Call(command, data, &reply)
	w.setState(worker, true)
	return
}
