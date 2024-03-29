package main

import (
	"fmt"
	"net/http"
)

func howdy(w http.ResponseWriter, req *http.Request) {
        fmt.Fprintf(w, "Howdy there!\n")
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hey there!\n")
}

func endpoint(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "This is a sweet endpoint, bro\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":6969", nil)
}
