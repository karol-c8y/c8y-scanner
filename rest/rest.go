package rest

import (
	"c8y-scanner/cumulocity"
	"c8y-scanner/scanner"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
)

type HttpHandler struct {
	m                  *cumulocity.Microservice
	filesToScanChannel *chan string
}

func (handler *HttpHandler) hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Try {url}/scan/{id} and watch for Events/Alarms")
}

func (handler *HttpHandler) scanBinaryId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var gid string = vars["id"]

	message := fmt.Sprintf("Scanning binary id %s", gid)
	handler.m.RaiseEvent("c8y-scanner", message)
	fmt.Fprintf(w, message)

	*handler.filesToScanChannel <- gid
}

func (handler *HttpHandler) scanFile(w http.ResponseWriter, r *http.Request) {
	tmpDir, _ := os.MkdirTemp("", "scan")
	tmpFile, _ := os.CreateTemp(tmpDir, "scan")
	io.Copy(tmpFile, r.Body)

	f := cumulocity.CleanableFile{Filename: tmpFile.Name()}
	defer f.Clean()

	res := scanner.Scan(f.Filename)
	if res.Vulnerable {
		fmt.Fprintln(w, "vulnerable", res.Description)
	} else {
		fmt.Fprintln(w, "not vulnerable")
	}
}

func Init(m *cumulocity.Microservice, filesToScanChannel *chan string) {
	handler := HttpHandler{m: m, filesToScanChannel: filesToScanChannel}
	r := mux.NewRouter()
	r.HandleFunc("/", handler.hello).Methods("GET")
	r.HandleFunc("/scan", handler.scanFile).Methods("POST")
	r.HandleFunc("/scan/{id}", handler.scanBinaryId).Methods("PUT")

	http.ListenAndServe(":80", r)
}
