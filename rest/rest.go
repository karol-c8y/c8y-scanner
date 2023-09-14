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
	m *cumulocity.Microservice
}

func (handler *HttpHandler) hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Try {url}/scan/{id} and watch for Events/Alarms")
}

func (handler *HttpHandler) scanBinaryId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gid := vars["id"]

	filepath := handler.m.DownloadFile(gid)
	defer handler.m.CleanupFile(filepath)

	message := fmt.Sprintf("Scanning binary id %s, downloaded filename: %s", gid, filepath)
	handler.m.RaiseEvent("c8y-scanner", message)
	fmt.Fprintf(w, message)
}

func (handler *HttpHandler) scanFile(w http.ResponseWriter, r *http.Request) {
	file, _ := os.CreateTemp("", "tmp")
	defer os.Remove(file.Name())
	io.Copy(file, r.Body)

	res := scanner.Scan(file.Name())
	if res.Vulnerable {
		fmt.Fprintln(w, "vulnerable", res.Description)
	} else {
		fmt.Fprintln(w, "not vulnerable")
	}
}

func Init(m *cumulocity.Microservice) {
	handler := HttpHandler{m: m}
	r := mux.NewRouter()
	r.HandleFunc("/", handler.hello).Methods("GET")
	r.HandleFunc("/scan", handler.scanFile).Methods("POST")
	r.HandleFunc("/scan/{id}", handler.scanBinaryId).Methods("PUT")

	http.ListenAndServe(":80", r)
}
