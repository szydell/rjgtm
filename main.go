package main

import "os"
import "fmt"
import "github.com/szydell/gogtm"

func main() {
	gtmGblDir := os.Getenv("gtmgbldir")
	fmt.Println("gtmgbldir: " + gtmGblDir)
	err := gogtm.Start()

}
