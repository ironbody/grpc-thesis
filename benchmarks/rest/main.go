package main

import (
	pb "Benchmarks/grpc/experiment"
	"encoding/json"
	"io"
	"log"
	"net/http"
)


var testItem  pb.Item = pb.Item{
	Id: "641eeb68f722093cd36746cd",
	Index: 0,
	Guid: "8c30a3df-7a8c-4c75-904d-6c8464db486a",
	IsActive: false,
	Balance: 1569.64,
	Picture: "http://placehold.it/32x32",
	Age: 22,
	EyeColor: "blue",
	Name: "Jeannie Hodge",	
	Gender: "female",
	Company: "ATGEN",
	Email: "jeanniehodge@atgen.com",
	Phone: "+1 (995) 400-2899",
	Address: "487 Beverley Road, Ahwahnee, Northern Mariana Islands, 1368",
	About: "Minim do nisi officia nisi nisi velit est. Labore est ex enim qui Lorem aliqua. Aliqua officia reprehenderit minim duis laborum amet occaecat.",
	Registered: "2020-06-27T08:59:26 -02:00",
	Latitude: 31.12956,
	Longitude: -10.251515,
	Greeting: "Hello, Jeannie Hodge! You have 10 unread messages.",
	FavoriteFruit: "strawberry",
}
var test2Return []*pb.Item
var test3Return []*pb.Item

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
	test2Return = make([]*pb.Item, 150)
	for i := 0; i < 150; i++ {
		test2Return[i] = &testItem
	}

	test3Return = make([]*pb.Item, 4500)
	for i := 0; i < 4500; i++ {
		test3Return[i] = &testItem
	}

	test2data,_ := json.Marshal(test2Return)
	test2Json = string(test2data)

	test3data,_ := json.Marshal(test3Return)
	test3Json = string(test3data)

	http.HandleFunc("/test1", Test1)
	http.HandleFunc("/test2", Test2)
	http.HandleFunc("/test3", Test3)

	http.ListenAndServe(":8080", nil)
}
