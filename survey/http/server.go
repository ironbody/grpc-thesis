package main

import (
	"net/http"
)

func retrieveUser(w http.ResponseWriter, req *http.Request) {

}

func main() {
	http.HandleFunc("/retrieveUser", retrieveUser)
	http.ListenAndServe(":9009", nil)
}
