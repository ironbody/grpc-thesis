package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello234\n")
}

var num = 32

var expectedUser = map[string]interface{}{
	"id":   float64(1969),
	"name": "Margaret Hamilton",
}

var expectedNumber = map[string]int{
	"number": 64,
}

func getNumber(w http.ResponseWriter, req *http.Request) {
	res := map[string]int{
		"number": num,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func checkNumber(w http.ResponseWriter, req *http.Request) {
	data := map[string]interface{}{}

	decoder := json.NewDecoder(req.Body)
	decoder.Decode(&data)

	log.Println(data)
	log.Println(expectedNumber)

	if fmt.Sprint(data) == fmt.Sprint(expectedNumber) {
		fmt.Fprintf(w, "Correct!")
	} else {
		fmt.Fprintf(w, "Incorrect! (The data is case sensitive)")
	}
}

func checkUser(w http.ResponseWriter, req *http.Request) {
	host, _, _ := net.SplitHostPort(req.RemoteAddr)
	addr := fmt.Sprintf("http://%s:9009/retrieveUser", host)

	resp, err := http.Get(addr)
	if err != nil {
		fmt.Fprintf(w, "Server error trying to contact %s", addr)
		log.Printf("Server error trying to contact %s", addr)
		return
	}

	data := map[string]interface{}{}

	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&data)

	log.Println(data)
	log.Println(expectedUser)

	if fmt.Sprint(data) == fmt.Sprint(expectedUser) {
		fmt.Fprintf(w, "Correct!")
	} else {
		fmt.Fprintf(w, "Incorrect! (The answer is case sensitive)")
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/getNumber", getNumber)
	http.HandleFunc("/checkNumber", checkNumber)
	http.HandleFunc("/checkUser", checkUser)

	http.ListenAndServe(":8080", nil)
}
