package scanner

import (
	"fmt"
	"github.com/dutchcoders/go-clamd"
	"os"
	"path"
	"time"
)

type ScanResult struct {
	Vulnerable  bool
	Description string
}

func Wait() {
	clam := clamd.NewClamd("/tmp/clamd.sock")
	err := clam.Ping()
	for err != nil {
		time.Sleep(2 * time.Second)
		err = clam.Ping()
	}
}

func Scan(filePath string) ScanResult {
	fmt.Printf("Scanning file %s\n", filePath)

	clam := clamd.NewClamd("/tmp/clamd.sock")

	allowFileRead(filePath)
	response, _ := clam.ContScanFile(filePath)

	for s := range response {
		//fmt.Printf("scan Raw: %v\n", s.Raw)
		//fmt.Printf("scan Description: %v\n", s.Description)
		//fmt.Printf("scan Hash: %v\n", s.Hash)
		//fmt.Printf("scan Path: %v\n", s.Path)
		//fmt.Printf("scan Size: %v\n", s.Size)
		//fmt.Printf("scan Status: %v\n", s.Status)
		//fmt.Printf("scan err %v\n", err)

		switch s.Status {
		case "FOUND":
			return ScanResult{Vulnerable: true, Description: s.Description}
		default:
			return ScanResult{}
		}
	}

	return ScanResult{}
}

func allowFileRead(filePath string) {
	dirPath := path.Dir(filePath)
	os.Chmod(dirPath, 0555)
	os.Chmod(filePath, 0444)
}
