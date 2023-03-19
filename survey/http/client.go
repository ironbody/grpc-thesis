package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func getModifiedNumber(w http.ResponseWriter, req *http.Request) {
	//res := map[string]int{
	//	"number": 64,
	//}
	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(res)

	data := fmt.Sprintf(`{“friends”:5,”favoriteThing”:”dogs”}`)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, data)
}

func retrieveUser(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, `{"Id":1969,"Name":"Margaret Hamilton"}`)
}

func main() {
	// go callServer()

	http.HandleFunc("/getModifiedNumber", getModifiedNumber)
	http.HandleFunc("/retrieveUser", retrieveUser)
	http.ListenAndServe(":9009", nil)
}

func callServer() {
	data := `{"number":64}`
	reader := bytes.NewReader([]byte(data))
	resp, err := http.Post("http://localhost:8080/checkNumber", "application/json", reader)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	println(string(body))
}
