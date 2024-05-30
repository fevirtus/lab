package main

import (
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadGateway)
	w.Write([]byte("Bad Gateway"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
	http.ListenAndServe(":443", nil)
	select {}
}
