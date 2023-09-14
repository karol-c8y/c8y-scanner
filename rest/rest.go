package rest

import (
	"fmt"
	"net/http"
	"c8y-scanner/cumulocity"

	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Try {url}/scan/{id} and watch for Events/Alarms")
}

func scanBinaryId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gid := vars["id"]

	filepath := cumulocity.DownloadFile(gid)

	message := fmt.Sprintf("Scanning binary id %s, downloaded filename: %s", gid, filepath)
	cumulocity.RaiseEvent("c8y-scanner", message)
	fmt.Fprintf(w, message)
}

func Init() {
	r := mux.NewRouter()
	r.HandleFunc("/", hello)
	r.HandleFunc("/scan/{id}", scanBinaryId)

	http.ListenAndServe(":80", r)
}
