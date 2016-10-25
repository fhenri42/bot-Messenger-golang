package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func route(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("we are in the route\n")
	var t string
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	if err != nil {
	fmt.Printf("in the error")
	fmt.Println(err)
	}
	fmt.Println("res" ,t)
	fmt.Printf("res" ,t)
}

func callRecast() {
fmt.Printf("i am in the function call recast\n")

}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/webhook", route)

	callRecast()
	http.ListenAndServe(":8080", nil)
}
