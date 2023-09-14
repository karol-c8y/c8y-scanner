package main

import (
	"c8y-scanner/cumulocity"
	"c8y-scanner/rest"
	"c8y-scanner/scanner"
)

func main() {
	cumulocity.Init()
	scanner.Wait()
	go rest.Init()
	scanner.Test()
}
