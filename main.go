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

		log.Println("List of already subscribed workers:")
		for key, value := range workers.returnWorkers() {
			fmt.Println("Worker:", key, "Value:", value)
		}
	})
	server.OnConnect(Subscribe)

	// type Args struct{ A, B int }
	// type Reply int

	// server.Handle("add", func(client *rpc2.Client, args *Args, reply *Reply) error {
	// 	// Reversed call (server to client)
	// 	var rep Reply
	// 	client.Call("mult", Args{2, 3}, &rep)
	// 	fmt.Println("mult result:", rep)

	// 	*reply = Reply(args.A + args.B)
	// 	return nil
	// })

	// server.Handle("dupadupa", func(client *rpc2.Client, args *Args, reply *Reply) error {
	// 	fmt.Println("Subscribing now...")
	// 	*reply = Reply(50)
	// 	return nil
	// })

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
	router.GET("/v1/data/:glvn", getGlvn)
	//Start listening for clients
	err := srv.ListenAndServe()
	mstools.ErrCheck(err)

}

//Subscribe function for workers
func Subscribe(worker *rpc2.Client) {
	log.Println("Subscribing new worker on connect...")
	workers.subscribeWorker(worker)
	log.Println("List of already subscribed workers:")
	for key, value := range workers.returnWorkers() {
		fmt.Println("Worker:", key, "Value:", value)
	}
}
