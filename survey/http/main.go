package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello234\n")
}

var num = 32

type NumberResponse struct {
	Number int
}

type UserResponse struct {
	Id   int
	Name string
}

var expectedUser = UserResponse{Id: 1969, Name: "Margaret Hamilton"}

var expectedNumber = NumberResponse{Number: num * 2}

func getNumber(w http.ResponseWriter, req *http.Request) {
	res := map[string]int{
		"number": num,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func checkNumber(w http.ResponseWriter, req *http.Request) {
	num := NumberResponse{}

	decoder := json.NewDecoder(req.Body)
	decoder.Decode(&num)

	log.Println(num.Number)
	log.Println(expectedNumber)

	if num.Number == expectedNumber.Number {
		fmt.Fprintf(w, "Correct!")
	} else {
		log.Printf("%d != %d\n", num.Number, expectedNumber.Number)
		fmt.Fprintf(w, "Incorrect! (The data is case sensitive)")
	}
}

func checkUser(w http.ResponseWriter, req *http.Request) {

	resp, err := http.Get("http://localhost:9009/retrieveUser")
	if err != nil {
		fmt.Fprintf(w, "server error trying to contact server")
		log.Printf("server error trying to contact server")
		return
	}

	var data UserResponse
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&data)

	log.Println(data)
	log.Println(expectedUser)

	if data == expectedUser {
		fmt.Fprintf(w, "Correct!")
	}

	fmt.Fprintf(w, "Incorrect! (The answer is case sensitive)")
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/getNumber", getNumber)
	http.HandleFunc("/checkNumber", checkNumber)
	http.HandleFunc("/checkUser", checkUser)
	http.ListenAndServe(":8080", nil)
}
