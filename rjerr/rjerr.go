package rjerr

import "errors"

//ErrNoAvailableWorkers means that no workers are available
var ErrNoAvailableWorkers = errors.New("No workers available")

//ErrAllWorkersBusy means that all workers are busy
var ErrAllWorkersBusy = errors.New("All workers busy")

//ErrGtmCantGetGlvn means that GT.M db did not respond properly after 'Get'
var ErrGtmCantGetGlvn = errors.New("GT.M error. Get does not respond properly")

//Err404 data not found
var Err404 = errors.New("Not found")
