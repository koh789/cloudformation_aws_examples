package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Println("receive request. ", req.RequestURI)
		io.WriteString(w, "OK")
	})
	http.HandleFunc("//system/health_check", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "OK")
	})
	log.Println("listen port:8080..")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
