package main

import (
	"flag"
	"fmt"
	"net/http"
)

const (
	port = ":3333"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func init() {
	if flag.Lookup("test.v") == nil {
		fmt.Printf("Started server at http://localhost%v.\n", port)
		http.HandleFunc("/", HelloWorld)
		http.ListenAndServe(port, nil)
	} else {
		return
	}
}

func main() {}
