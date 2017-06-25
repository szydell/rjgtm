package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func getGlvn(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	glvn := string(ps.ByName("glvn"))
	log.Println("GET glvn:" + glvn)
	reply, err := workers.doWork("getGlvn", glvn)
	if err == errAllWorkersBusy {
		http.Error(w, "504 Timeout", 504)
	}
	if err == errNoAvailableWorkers {
		http.Error(w, "503 Service Unavailable", 503)
	}
	log.Println(reply, err)
}
