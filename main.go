package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/cenkalti/rpc2"
	"github.com/julienschmidt/httprouter"
	"github.com/szydell/mstools"
)

var server = rpc2.NewServer()

var workers = newAllWorkers()

func main() {

	server.OnDisconnect(func(worker *rpc2.Client) {
		workers.unSubscribeWorker(worker)

		showWorkers()
	})
	server.OnConnect(Subscribe)

	//Start listening for workers
	lis, _ := net.Listen("tcp", "127.0.0.1:5000")
	go server.Accept(lis)
	defer lis.Close()

	//define a http server for RESTful queries
	router := httprouter.New()
	addr := ":8080"

	// prepare server
	srv := &http.Server{Addr: addr, Handler: router}

	//define routes
	router.GET("/v1/global/:glvn", getGlvn)
	router.POST("/v1/global/:glvn", setGlvn)
	router.GET("/v1/gvstats", getGvStat)
	router.DELETE("/v1/gvstats", deleteGvStat)
	router.GET("/v1/order/:glvn", orderGlvn)
	router.GET("/v1/prev/:glvn", prevGlvn)
	router.GET("/v1/query/:glvn", queryGlvn)
	router.GET("/v1/data/:glvn", dataGlvn)

	//Start listening for clients
	err := srv.ListenAndServe()
	mstools.ErrCheck(err)

}

//Subscribe function for workers
func Subscribe(worker *rpc2.Client) {
	log.Println("Subscribing new worker on connect...")
	workers.subscribeWorker(worker)
	showWorkers()
}

func showWorkers() {
	log.Println("List of already subscribed workers:")
	for key, value := range workers.returnWorkers() {
		fmt.Println("Worker:", key, "Value:", value)
	}
}
