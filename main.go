package main

import (
	"c8y-scanner/cumulocity"
	"c8y-scanner/rest"
	"c8y-scanner/scanner"
)

func main() {
	ms := cumulocity.Init()
	go rest.Init(&ms)
	scanner.Test()
}
