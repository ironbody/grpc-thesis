package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

var test2Json string
var test3Json string


func Test1(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, World!")
}

func Test2(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, test2Json)
}

func Test3(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, test3Json)
}


func main() {
	log.Println("HTTP Server Started...")

	dat1, _ := os.ReadFile("150k.txt")
	dat2, _ := os.ReadFile("3M.txt")

	test2Json = string(dat1)
	test3Json = string(dat2)

	http.HandleFunc("/test1", Test1)
	http.HandleFunc("/test2", Test2)
	http.HandleFunc("/test3", Test3)

	http.ListenAndServe(":8080", nil)
}
