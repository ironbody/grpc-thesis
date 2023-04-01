package main

import (
	"fmt"
	_ "io/ioutil"
	_ "log"
	"net/http"
	_ "net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func main() {
	http.HandleFunc("/test1", hello)

	http.ListenAndServe(":8080", nil)
}
