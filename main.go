package main

import (
	"c8y-scanner/rest"
	"c8y-scanner/scanner"
	"c8y-scanner/cumulocity"
)

func main() {
	cumulocity.Init()
	go rest.Init()
	scanner.Scan()
}
