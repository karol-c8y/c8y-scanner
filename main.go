package main

import (
	"c8y-scanner/rest"
	"c8y-scanner/scanner"
)

func main() {
	go rest.Init()
	scanner.Scan()
}
