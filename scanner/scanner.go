package scanner

import (
	"bytes"
	"fmt"
	"github.com/dutchcoders/go-clamd"
	"time"
)

func Scan() {
	fmt.Println("Made with <3 DutchCoders")

	c := clamd.NewClamd("/tmp/clamd.sock")
	//_ = c

	reader := bytes.NewReader(clamd.EICAR)
	response, err := c.ScanStream(reader, make(chan bool))
	for s := range response {
		fmt.Printf("scan stream: %v %v\n", s, err)
	}

	response, err = c.AllMatchScanFile("./testfiles/")
	for s := range response {
		fmt.Printf("scan file: %v %v\n", s, err)
	}

	response, err = c.Version()
	for s := range response {
		fmt.Printf("version: %v %v\n", s, err)
	}

	for {
		err := c.Ping()
		fmt.Printf("Ping: %v\n", err)

		stats, err := c.Stats()
		fmt.Printf("%v %v\n", stats, err)

		// err = c.Reload()
		// fmt.Printf("Reload: %v\n", err)

		time.Sleep(2 * time.Second)
	}
	// response, err = c.Shutdown()
	// fmt.Println(response)
}
