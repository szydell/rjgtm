package main

import "os"
import "fmt"
import "github.com/szydell/gogtm"
import "github.com/szydell/mstools"

func main() {
	gtmGblDir := os.Getenv("gtmgbldir")
	fmt.Println("gtmgbldir: " + gtmGblDir)
	gtmDist := os.Getenv("gtm_dist")
	fmt.Println("gtm_dist: " + gtmDist)
	err := gogtm.Start()
	mstools.ErrCheck(err)

}
