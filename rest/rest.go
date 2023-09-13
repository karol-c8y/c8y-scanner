package rest

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func Init() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homePage)

	http.ListenAndServe(":8888", mux)
}
