package rjerr

import (
	"errors"
	"strings"
)

//ErrNoAvailableWorkers means that no workers are available
var ErrNoAvailableWorkers = errors.New("No workers available")

//ErrAllWorkersBusy means that all workers are busy
var ErrAllWorkersBusy = errors.New("All workers busy")

//ErrGtmCantGetGlvn means that GT.M db did not respond properly after 'Get'
var ErrGtmCantGetGlvn = errors.New("GT.M error. Get does not respond properly")

//Err404 data not found
var Err404 = errors.New("Not found")

//ErrorTypeAndMessage try to figure out what type of http error should be returned
func ErrorTypeAndMessage(reply string) (id int) {

	switch {
	case strings.Contains(reply, "{\"ERROR"):
		id = 503
	case strings.Contains(reply, "not found"):
		id = 404
	default:
		id = 500
	}
	return
}
