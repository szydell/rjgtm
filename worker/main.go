package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"errors"

	"github.com/cenkalti/rpc2"
	uuid "github.com/satori/go.uuid"
	"github.com/szydell/gogtm"
	"github.com/szydell/mstools"
)

var id = uuid.NewV4().String()

func main() {
	gtmGblDir := os.Getenv("gtmgbldir")
	fmt.Println("gtmgbldir: " + gtmGblDir)
	gtmDist := os.Getenv("gtm_dist")
	fmt.Println("gtm_dist: " + gtmDist)
	err := gogtm.Start()
	defer gogtm.Stop()
	mstools.ErrCheck(err)

	conn, _ := net.Dial("tcp", "127.0.0.1:5000")
	client := rpc2.NewClient(conn)
	client.Handle("getGlvn", getGlvn)
	client.Run()

}

func getGlvn(client *rpc2.Client, glvn string, reply *interface{}) error {

	log.Println("GET glvn:" + glvn)
	response, err := gogtm.Get("^"+glvn, id)
	if err != nil {
		log.Println("503 /v1/data/" + glvn)
		*reply = nil
		return errors.New("Błąd przy pobieraniu globala")
	}
	if response == id {
		log.Println("404 /v1/data/" + glvn)
		*reply = nil
		return errors.New("404")
	}
	//return string formatted as JSON, try to figure out if response is a string or integer
	if _, err := strconv.Atoi(response); err == nil {
		*reply = "{\"" + glvn + "\": " + response + "}"
	} else {
		*reply = "{\"" + glvn + "\": \"" + response + "\"}"
	}
	log.Println("200 /v1/data/" + glvn + " -> " + response)

	return nil
}
