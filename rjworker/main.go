package main

import (
	"log"
	"net"
	"strconv"

	"github.com/cenkalti/rpc2"
	uuid "github.com/satori/go.uuid"
	"github.com/szydell/gogtm"
	"github.com/szydell/mstools"
	"github.com/szydell/rjgtm/rjerr"
	"github.com/szydell/rjgtm/rjshared"
)

var id = uuid.NewV4().String()

func main() {

	err := gogtm.Start()
	defer gogtm.Stop()
	mstools.ErrCheck(err)

	conn, _ := net.Dial("tcp", "127.0.0.1:5000")
	client := rpc2.NewClient(conn)
	client.Handle("getGlvn", getGlvn)
	client.Handle("gvStats", gvStats)
	client.Handle("cleanGvStats", cleanGvStats)
	client.Handle("setGlvn", setGlvn)
	client.Handle("orderGlvn", orderGlvn)
	client.Handle("prevGlvn", prevGlvn)
	client.Handle("queryGlvn", queryGlvn)
	client.Run()

}

func getGlvn(client *rpc2.Client, glvn string, reply *string) error {

	log.Println("GET glvn:" + glvn)
	response, err := gogtm.Get("^"+glvn, id)
	if err != nil {
		log.Println("503 /v1/data/" + glvn)
		return rjerr.ErrGtmCantGetGlvn
	}
	if response == id {
		log.Println("404 /v1/data/" + glvn)
		return rjerr.Err404
	}
	//return string formatted as JSON, try to figure out if response is a string or integer
	if _, err := strconv.Atoi(response); err == nil {
		*reply = "{\"RESPONSE\":{\"" + glvn + "\": " + response + "}, \"STATUS\":\"OK\"}"
	} else {
		*reply = "{\"RESPONSE\":{\"" + glvn + "\": \"" + response + "\"}, \"STATUS\":\"OK\"}"
	}
	log.Println("200 /v1/data/" + glvn + " -> " + response)

	return nil
}

func gvStats(client *rpc2.Client, _, reply *string) error {

	log.Println("GET gvstat")
	response, err := gogtm.GvStat()
	if err != nil {
		log.Println("503 /v1/gvstat")
		return rjerr.ErrGtmCantGetGlvn
	}

	buildJSON := []rune("{\"RESPONSE\":[{\"")

	for _, char := range response {
		switch char {
		case 44:
			buildJSON = append(buildJSON, []rune{',', '"'}...)
		case 58:
			buildJSON = append(buildJSON, []rune{'"', ':'}...)
		case 59:
			buildJSON = append(buildJSON, []rune{'"', ':', '[', '{', '"'}...)
		case 124:
			buildJSON = append(buildJSON, []rune{'}', ']', ',', '"'}...)
		default:
			buildJSON = append(buildJSON, rune(char))
		}
	}
	buildJSON = append(buildJSON, []rune("}]}],\"STATUS\":\"OK\"}")...)
	*reply = string(buildJSON)
	return nil
}

func cleanGvStats(client *rpc2.Client, _, reply *string) error {
	log.Println("DELETE gvstat")
	err := gogtm.Xecute("S REGION=$V(\"GVFIRST\") VIEW \"GVSRESET\":REGION F I=1:1 S REGION=$V(\"GVNEXT\",REGION) Q:REGION=\"\"  VIEW \"GVSRESET\":REGION")
	if err == nil {
		*reply = "{\"status\":\"OK\"}"
	} else {
		log.Println("503 DELETE gvstats failed (cleanGvStats function)")
		*reply = "{\"status\":\"ERROR\"}"
	}
	return err
}

func setGlvn(client *rpc2.Client, glvn rjshared.Glvn, reply *string) error {
	log.Println("POST glvn:" + glvn.Key + "(setGlvn function)")
	err := gogtm.Set("^"+glvn.Key, glvn.Value)

	if err != nil {
		log.Println("503 SET /v1/data/" + glvn.Key)
		return rjerr.ErrGtmCantSetGlvn
	}
	*reply = "{\"status\":\"OK\"}"

	return nil
}

func orderGlvn(client *rpc2.Client, glvn string, reply *string) error {

	log.Println("ORDER glvn:" + glvn)
	response, err := gogtm.Order("^"+glvn, "")
	if err != nil {
		log.Println("503 /v1/order/" + glvn)
		return rjerr.ErrGtmCantGetGlvn
	}

	//return string formatted as JSON, try to figure out if response is a string or integer
	if _, err := strconv.Atoi(response); err == nil {
		*reply = "{\"RESPONSE\":{\"next glvn\": " + response + "}, \"STATUS\":\"OK\"}"
	} else {
		*reply = "{\"RESPONSE\":{\"next glvn\": \"" + response + "\"}, \"STATUS\":\"OK\"}"
	}
	log.Println("200 /v1/order/" + glvn + " -> next glvn:" + response)

	return nil
}

func prevGlvn(client *rpc2.Client, glvn string, reply *string) error {

	log.Println("PREVIOUS glvn:" + glvn)
	response, err := gogtm.Order("^"+glvn, "-1")
	if err != nil {
		log.Println("503 /v1/prev/" + glvn)
		return rjerr.ErrGtmCantGetGlvn
	}

	//return string formatted as JSON, try to figure out if response is a string or integer
	if _, err := strconv.Atoi(response); err == nil {
		*reply = "{\"RESPONSE\":{\"previous glvn\": " + response + "}, \"STATUS\":\"OK\"}"
	} else {
		*reply = "{\"RESPONSE\":{\"previous glvn\": \"" + response + "\"}, \"STATUS\":\"OK\"}"
	}
	log.Println("200 /v1/prev/" + glvn + " -> previous glvn:" + response)

	return nil
}

func queryGlvn(client *rpc2.Client, glvn string, reply *string) error {

	log.Println("QUERY glvn:" + glvn)
	response, err := gogtm.Query("^" + glvn)
	if err != nil {
		log.Println("503 /v1/query/" + glvn)
		return err
	}

	//return string formatted as JSON, try to figure out if response is a string or integer
	if _, err := strconv.Atoi(response); err == nil {
		*reply = "{\"RESPONSE\":{\"query result\": " + response + "}, \"STATUS\":\"OK\"}"
	} else {
		*reply = "{\"RESPONSE\":{\"query result\": \"" + response + "\"}, \"STATUS\":\"OK\"}"
	}
	log.Println("200 /v1/order/" + glvn + " -> query result:" + response)

	return nil
}
