package main

import (
	"net/http"
	"strconv"

	"log"

	"fmt"

	"github.com/julienschmidt/httprouter"
	"github.com/szydell/gogtm"
)

func getGlvn(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	glvn := string(ps.ByName("glvn"))
	log.Println("GET glvn:" + glvn)
	response, err := gogtm.Get("^"+glvn, suuid)
	if err != nil {
		http.Error(w, "503 Service Unavailable", 503)
		log.Println("503 /v1/data/" + glvn)
		return
	}
	if response == suuid {
		http.Error(w, "404 Not Found", 404)
		log.Println("404 /v1/data/" + glvn)
		return
	}
	//return string formatted as JSON, try to figure out if response is a string or integer
	if _, err := strconv.Atoi(response); err == nil {
		fmt.Fprintf(w, "{\"%s\": %s}", glvn, response)
	} else {
		fmt.Fprintf(w, "{\"%s\": \"%s\"}", glvn, response)
	}
	log.Println("200 /v1/data/" + glvn + " -> " + response)
}

func count(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	glvn := string(ps.ByName("glvn"))
	log.Println("COUNT, glvn:" + glvn)
	response, err := gogtm.Xecute("")
	if err != nil {
		http.Error(w, "503 Service Unavailable", 503)
		log.Println("503 /v1/data/" + glvn + "/count")
		return
	}
	log.Println("Count:" + response)
}

func halt(srv *http.Server) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		gogtm.Stop()
		srv.Shutdown(nil)
	}
}
