package main

import (
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
	fmt.Printf("Started server at http://localhost%v.\n", port)
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(port, nil)
}

func main() {}
