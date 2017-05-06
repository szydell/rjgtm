package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"
	"github.com/szydell/gogtm"
	"github.com/szydell/mstools"
)

//session global uuid, generated on start
var suuid string

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go handleCtrlC(c)

	gtmGblDir := os.Getenv("gtmgbldir")
	fmt.Println("gtmgbldir: " + gtmGblDir)
	gtmDist := os.Getenv("gtm_dist")
	fmt.Println("gtm_dist: " + gtmDist)
	err := gogtm.Start()
	defer gogtm.Stop()
	mstools.ErrCheck(err)
	// generate session global uid for multiple purposes
	suuid = uuid.NewV4().String()

	// we are connected to the database, so prepare router!

	router := httprouter.New()
	addr := ":8080"

	// prepare server
	srv := &http.Server{Addr: addr, Handler: router}

	//define routes
	router.GET("/v1/data/:glvn", getGlvn)
	router.GET("/v1/ops/halt", halt(srv))

	//Start to listen
	err = srv.ListenAndServe()
	mstools.ErrCheck(err)
}
