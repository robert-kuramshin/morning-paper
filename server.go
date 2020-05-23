package main

import (
	"io"
	"net/http"
	"os/exec"
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
			res := ""
			for _, url := range feed_urls {
				res += parse_feed(url, 3)
			}
			c1 := exec.Command("echo", "-e", res)
			c2 := exec.Command("lpr")

			rP, wP := io.Pipe()
			c1.Stdout = wP
			c2.Stdin = rP

			c1.Start()
			c2.Start()
			c1.Wait()
			wP.Close()
			c2.Wait()
		}
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}
}

func main() {
	http.HandleFunc(apiPath, handleRequest)
	http.ListenAndServe(":80", nil)
}
