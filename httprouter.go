package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/szydell/mstools"
	"github.com/szydell/rjgtm/rjerr"
	"github.com/szydell/rjgtm/rjshared"
)

func getGlvn(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	glvn := string(ps.ByName("glvn"))
	log.Println("GET glvn:" + glvn)
	w.Header().Set("Content-Type", "application/json")
	reply, err := workers.doWork("getGlvn", glvn)
	if err != nil {
		tmpID, tmpDescr := rjerr.ErrorTypeAndMessage(err)
		log.Println(tmpID, tmpDescr)
		http.Error(w, tmpDescr, tmpID)
		return
	}

	fmt.Fprintf(w, "%s", reply)
	log.Println("200 ", reply, err)
}

func getGvStat(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	log.Println("GET gvstat")
	w.Header().Set("Content-Type", "application/json")
	reply, err := workers.doWork("gvStats", "")
	if err != nil {
		tmpID, tmpDescr := rjerr.ErrorTypeAndMessage(err)
		log.Println(tmpID, tmpDescr)
		http.Error(w, tmpDescr, tmpID)
		return
	}
	fmt.Fprintf(w, "%s", reply)
	log.Println("200", reply, err)
}

func deleteGvStat(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Println("DELETE gvstat")
	w.Header().Set("Content-Type", "application/json")
	reply, err := workers.doWork("cleanGvStats", "")
	if err != nil {
		tmpID, tmpDescr := rjerr.ErrorTypeAndMessage(err)
		log.Println(tmpID, tmpDescr)
		http.Error(w, tmpDescr, tmpID)
		return
	}
	fmt.Fprintf(w, "%s", reply)
	log.Println("200", reply, err)
}

func setGlvn(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var glvn rjshared.Glvn
	glvn.Key = ps.ByName("glvn")
	log.Println("SET ", glvn.Key)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		mstools.ErrCheck(err)
	}
	glvn.Value = string(body)
	w.Header().Set("Content-Type", "application/json")
	reply, err := workers.doWork("setGlvn", glvn)
	if err != nil {
		tmpID, tmpDescr := rjerr.ErrorTypeAndMessage(err)
		log.Println(tmpID, tmpDescr)
		http.Error(w, tmpDescr, tmpID)
		return
	}
	fmt.Fprintf(w, "%s", reply)
	log.Println("200", reply, err)
}

func orderGlvn(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	glvn := string(ps.ByName("glvn"))
	log.Println("ORDER glvn:" + glvn)
	w.Header().Set("Content-Type", "application/json")
	reply, err := workers.doWork("orderGlvn", glvn)
	if err != nil {
		tmpID, tmpDescr := rjerr.ErrorTypeAndMessage(err)
		log.Println(tmpID, tmpDescr)
		http.Error(w, tmpDescr, tmpID)
		return
	}

	fmt.Fprintf(w, "%s", reply)
	log.Println("200 ", reply, err)
}

func prevGlvn(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	glvn := string(ps.ByName("glvn"))
	log.Println("PREVIOUS glvn:" + glvn)
	w.Header().Set("Content-Type", "application/json")
	reply, err := workers.doWork("prevGlvn", glvn)
	if err != nil {
		tmpID, tmpDescr := rjerr.ErrorTypeAndMessage(err)
		log.Println(tmpID, tmpDescr)
		http.Error(w, tmpDescr, tmpID)
		return
	}

	fmt.Fprintf(w, "%s", reply)
	log.Println("200 ", reply, err)
}
