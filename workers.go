package main

import (
	"sync"

	"github.com/cenkalti/rpc2"
)

type allWorkers struct {
	sync.RWMutex
	list map[*rpc2.Client]bool
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

func (w *allWorkers) doWork(command string, data interface{}) (reply interface{}, err error) {
	if len(w.list) < 1 {
		return nil, errNoAvailableWorkers
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
		return nil, errAllWorkersBusy
	}
	w.list[worker] = false
	w.Unlock()
	err = worker.Call("getGlvn", data, &reply)
	w.setState(worker, true)
	return
}
