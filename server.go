package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var apiPath string = "/RPI/printer"

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != apiPath {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case "GET":
		for k, v := range r.URL.Query() {
			fmt.Printf("%s: %s\n", k, v)
		}
		w.Write([]byte("Recieved a GET request\n"))
	case "POST":
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", reqBody)
		w.Write([]byte("Recieved a POST request\n"))
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}
}

func main() {
	http.HandleFunc(apiPath, handleRequest)
	http.ListenAndServe(":80", nil)
}
