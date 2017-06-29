package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/szydell/rjgtm/rjerr"
)

func getGlvn(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	glvn := string(ps.ByName("glvn"))
	log.Println("GET glvn:" + glvn)
	w.Header().Set("Content-Type", "application/json")
	reply, err := workers.doWork("getGlvn", glvn)
	if err != nil {
		tmpID := rjerr.ErrorTypeAndMessage(reply)
		log.Println(tmpID, reply)
		http.Error(w, reply, tmpID)
		return
	}

	fmt.Fprintf(w, "%s", reply)
	log.Println(reply, err)
}

func getGvStat(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	log.Println("GET gvstat")
	w.Header().Set("Content-Type", "application/json")
	reply, err := workers.doWork("GvStats", "")
	if err != nil {
		tmpID := rjerr.ErrorTypeAndMessage(reply)
		log.Println(tmpID, reply)
		http.Error(w, reply, tmpID)
		return
	}
	fmt.Fprintf(w, "%s", reply)
	log.Println(reply, err)
}

func deleteGvStat(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Println("DELETE gvstat")
	w.Header().Set("Content-Type", "application/json")
	reply, err := workers.doWork("cleanGvStats", "")
	if err != nil {
		tmpID := rjerr.ErrorTypeAndMessage(reply)
		log.Println(tmpID, reply)
		http.Error(w, reply, tmpID)
		return
	}
	fmt.Fprintf(w, "%s", reply)
	log.Println(reply, err)
}
