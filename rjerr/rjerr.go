//Package rjerr contains errors shared between broker and workers
package rjerr

import (
	"errors"
)

//ErrNoAvailableWorkers means that no workers are available
var ErrNoAvailableWorkers = errors.New("No workers available")

//ErrAllWorkersBusy means that all workers are busy
var ErrAllWorkersBusy = errors.New("All workers busy")

//ErrGtmCantGetGlvn means that GT.M db did not respond properly after 'Get'
var ErrGtmCantGetGlvn = errors.New("GT.M error. Get does not respond properly")

//ErrGtmCantSetGlvn means that GT.M db did not respond properly after 'Set'
var ErrGtmCantSetGlvn = errors.New("GT.M error. Set could not be processed")

//Err404 data not found
var Err404 = errors.New("Not found")

//ErrorTypeAndMessage try to figure out what type of http error should be returned
func ErrorTypeAndMessage(err error) (id int, descr string) {

	switch err.Error() {
	case Err404.Error():
		id = 404
	case ErrNoAvailableWorkers.Error():
		id = 503
	case ErrGtmCantSetGlvn.Error():
		id = 503
	case ErrAllWorkersBusy.Error():
		id = 503
	default:
		id = 500
	}

	descr = "{\"STATUS\":\"" + err.Error() + "\"}"
	return
}
