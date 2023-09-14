package main

import (
	"c8y-scanner/cumulocity"
	"c8y-scanner/rest"
	"c8y-scanner/scanner"
	"fmt"
	"path"
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
	message_type := fmt.Sprintf("c8y-scanner-%s", gid)
	if result.Vulnerable {
		message := fmt.Sprintf("File %s (id: %s) is infected! %s", path.Base(file.Filename), gid, result.Description)
		m.RaiseMajorAlarm(message_type, message)
	} else {
		message := fmt.Sprintf("No vulnerability found in %s", file.Filename)
		m.RaiseEvent(message_type, message)
	}
}
