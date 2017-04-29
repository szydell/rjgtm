package main

import (
	"fmt"
	"os"

	uuid "github.com/satori/go.uuid"
	"github.com/szydell/gogtm"
	"github.com/szydell/mstools"
)

//session global uuid, generated on start
var suuid string

func main() {
	gtmGblDir := os.Getenv("gtmgbldir")
	fmt.Println("gtmgbldir: " + gtmGblDir)
	gtmDist := os.Getenv("gtm_dist")
	fmt.Println("gtm_dist: " + gtmDist)
	err := gogtm.Start()
	defer gogtm.Stop()
	mstools.ErrCheck(err)
	// generate session global uid for multiple purposes
	suuid = uuid.NewV4().String()

	//
	go cleanup()

	// we are connected to the database, so start listening!
	startRouter()

}
