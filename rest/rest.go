package rest

import (
	"c8y-scanner/cumulocity"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
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

func Init(m *cumulocity.Microservice) {
	handler := HttpHandler{m: m}
	r := mux.NewRouter()
	r.HandleFunc("/", handler.hello)
	r.HandleFunc("/scan/{id}", handler.scanBinaryId)

	http.ListenAndServe(":80", r)
}
