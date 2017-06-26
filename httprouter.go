package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/julienschmidt/httprouter"
	rjerr "github.com/szydell/rjgtm/rjerr"
)

func getGlvn(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	glvn := string(ps.ByName("glvn"))
	log.Println("GET glvn:" + glvn)
	reply, err := workers.doWork("getGlvn", glvn)
	if err != nil {
		log.Println(err)
		log.Println(reflect.TypeOf(err))
		tmpID, tmpDescr := errorTypeAndMessage(err)
		log.Println(tmpID, tmpDescr)
		http.Error(w, tmpDescr, tmpID)
		return
	}

	fmt.Fprintf(w, "%s", reply)
	log.Println(reply, err)
}

func errorTypeAndMessage(gotErr error) (errID int, errDescr string) {

	switch gotErr {
	case rjerr.Err404:
		errID = 404
		errDescr = gotErr.Error()
	case rjerr.ErrAllWorkersBusy:
		errID = 504
		errDescr = gotErr.Error()
	case rjerr.ErrNoAvailableWorkers, rjerr.ErrGtmCantGetGlvn:
		errID = 503
		errDescr = gotErr.Error()
	default:
		errID = 0
		errDescr = "unknown error"
	}
	return
}
