package main

import (
	"fmt"
	"os"
)

//I try to do some cleanup when got the signal to stop, but it is not working properly
func handleCtrlC(c chan os.Signal) {
	sig := <-c
	// handle ctrl+c event here
	// for example, close database
	fmt.Println("\nsignal: ", sig)
	os.Exit(0)
}
