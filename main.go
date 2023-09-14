package main

import (
	"c8y-scanner/cumulocity"
	"c8y-scanner/rest"
	"c8y-scanner/scanner"
	"fmt"
)

func main() {
	m := cumulocity.Init()
	filesToScanChannel := make(chan string, 100)
	go rest.Init(&m, &filesToScanChannel)
	scanner.Wait()
	for {
		var gid string = <-filesToScanChannel
		checkGid(&m, gid)
	}
}

func checkGid(m *cumulocity.Microservice, gid string) {
	file := m.DownloadFile(gid)
	defer file.Clean()

	fmt.Printf("Scanning file %s", file.Filename)

	result := scanner.Scan(file.Filename)
	if result.Vulnerable {
		message := fmt.Sprintf("%s is infected! %s", file.Filename, result.Description)
		m.RaiseMajorAlarm("c8y-scanner", message)
	} else {
		message := fmt.Sprintf("No vulnerability found in %s", file.Filename)
		m.RaiseEvent("c8y-scanner", message)
	}
}
