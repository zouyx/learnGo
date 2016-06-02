package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", http.HandlerFunc(handler)))
}

func handler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rw, "url.path: %s \n", req.URL.Path)
}
