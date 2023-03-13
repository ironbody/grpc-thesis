package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello234\n")
}

var num = 32

func getNumber(w http.ResponseWriter, req *http.Request) {
	res := map[string]int{
		"number": num,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func checkNumber(w http.ResponseWriter, req *http.Request) {
	println("check number called")
	fmt.Fprintf(w, "calling /getModifiedNumber in 3 seconds")

	go callClientNumber(req)
}

func callClientNumber(req *http.Request) {
	time.Sleep(3 * time.Second)

	host, _, _ := net.SplitHostPort(req.RemoteAddr)
	addr := fmt.Sprintf("http://%s:9999/getModifiedNumber", host)

	fmt.Printf("%s\n", addr)
	resp, err := http.Get(addr)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	println(string(body))
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/getNumber", getNumber)
	http.HandleFunc("/checkNumber", checkNumber)

	http.ListenAndServe(":8080", nil)
}
