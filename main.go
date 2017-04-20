package main

import "os"
import "fmt"

func main() {
	gtmGblDir := os.Getenv("gtmgbldir")
	fmt.Println("gtmgbldir: " + gtmGblDir)
}
