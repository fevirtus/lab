package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request headers: %v", r.Header)
	log.Printf("Request body: %v", r.Body)
	log.Printf("Request method: %v", r.Method)
	log.Printf("Request URL: %v", r.URL)
	log.Printf("Request host: %v", r.Host)
	log.Printf("Request remote address: %v", r.RemoteAddr)

	w.WriteHeader(http.StatusBadGateway)
	w.Write([]byte("Bad Gateway"))
}

func main() {
	http.HandleFunc("/", handler)

	go func() {
		err := http.ListenAndServe(":80", nil)
		if err != nil {
			log.Fatalf("HTTP server on port 80 failed: %s", err)
		}
	}()

	go func() {
		err := http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)
		if err != nil {
			log.Fatalf("HTTPS server on port 443 failed: %s", err)
		}
	}()

	log.Println("Server started")
	select {}
}
