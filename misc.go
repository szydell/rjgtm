package main

import (
	"os"
	"os/signal"

	"log"

	"github.com/szydell/gogtm"
)

//I try to do some cleanup when got the signal to stop, but it is not working properly
func cleanup() {

	log.Println("Cleanup possible.")
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)
	s := <-sigchan
	log.Println("got signal: ", s)
	gogtm.Stop()
	os.Exit(0)
}
