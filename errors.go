package main

import "errors"

var errNoAvailableWorkers = errors.New("No workers available")
var errAllWorkersBusy = errors.New("All workers busy")
