package main

import (
	"fmt"
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
		return
	}
	if err == errNoAvailableWorkers {
		http.Error(w, "503 Service Unavailable", 503)
		return
	}
	fmt.Fprintf(w, "%s", reply)
	log.Println(reply, err)
}
