package main

import (
	"fmt"
	"net/http"
)

var apiPath string = "/RPI/printer"
var apiPass string = "secret"

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != apiPath {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case "GET":
		correctPass := false

		for k, v := range r.URL.Query() {
			if k == "password" && v[0] == apiPass {
				correctPass = true
			}
		}
		w.Write([]byte("Recieved a GET request\n"))

		if correctPass {
			fmt.Println("Success")
		}
	// case "POST":
	// 	reqBody, err := ioutil.ReadAll(r.Body)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Printf("%s\n", reqBody)
	// 	w.Write([]byte("Recieved a POST request\n"))
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}
}

func main() {
	http.HandleFunc(apiPath, handleRequest)
	http.ListenAndServe(":80", nil)
}
