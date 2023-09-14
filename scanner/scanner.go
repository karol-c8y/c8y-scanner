package scanner

import (
	"fmt"
	"github.com/dutchcoders/go-clamd"
	"io"
	"net/http"
	"os"
	"time"
)

type ScanResult struct {
	Vulnerable  bool
	Description string
}

func Test() {
	clam := clamd.NewClamd("/tmp/clamd.sock")
	err := clam.Ping()
	for err != nil {
		time.Sleep(2 * time.Second)
		err = clam.Ping()
	}

	scan("https://www.google.com/robots.txt")
	scan("https://secure.eicar.org/eicar.com")
}

func scan(url string) {
	file, _ := os.CreateTemp("", "tmp")
	defer os.Remove(file.Name())

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	io.Copy(file, resp.Body)

	res := Scan(file.Name())
	if res.Vulnerable {
		fmt.Printf("%s vulnerable = %s\n", url, res.Description)
	} else {
		fmt.Printf("%s not vulnerable\n", url)
	}
}

func Scan(path string) ScanResult {
	fmt.Printf("Scan start file=%s\n", path)

	clam := clamd.NewClamd("/tmp/clamd.sock")

	os.Chmod(path, 444)
	response, _ := clam.ContScanFile(path)

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
