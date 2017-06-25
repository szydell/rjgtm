package main

import (
	"net"

	"github.com/cenkalti/rpc2"
	uuid "github.com/satori/go.uuid"
)

var id = uuid.NewV4().String()

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:5000")
	client := rpc2.NewClient(conn)
	client.Handle("getGlvn", getGlvn)
	client.Run()

}

func getGlvn(client *rpc2.Client, glvn string, reply *interface{}) error {
	*reply = "global wysłany przez workera " + id
	return nil
}
