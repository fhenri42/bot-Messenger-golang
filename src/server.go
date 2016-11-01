package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type Data struct {
	Entry []struct {
		ID        int64 `json:"id,string"`
		Messaging []struct {
			Message struct {
				Mid  string `json:"mid"`
				Seq  int64  `json:"seq"`
				Text string `json:"text"`
			} `json:"message"`
			Recipient struct {
				ID int64 `json:"id,string"`
			} `json:"recipient"`
			Sender struct {
				ID int64 `json:"id,string"`
			} `json:"sender"`
			Timestamp int64 `json:"timestamp"`
		} `json:"messaging"`
		Time int64 `json:"time"`
	} `json:"entry"`
	Object string `json:"object"`
}

func route(w http.ResponseWriter, req *http.Request) {

	if req.FormValue("hub.mode") == "subscribe" {
		token := req.FormValue("hub.verify_token")
		mode := req.FormValue("hub.mode")
		challenge := req.FormValue("hub.challenge")

		if (token == "123" && mode == "subscribe") {
			w.Header().Set("Server", "A Go Web Server")
			w.WriteHeader(200)
			w.Write([]byte(challenge))
			return
		} else {
			fmt.Printf("Your toke in not valid")
		}
	} else {
		fmt.Printf("in the post\n")
		fmt.Printf("%s\n", req)
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Println(err)
			return
		}
		var data Data
		err = json.Unmarshal(body, &data)
		if err != nil {
			log.Println(err)
			return
		}
		message_handler(data)
	}
	w.WriteHeader(200)
}

func main() {
	http.HandleFunc("/webhook", route)
	http.ListenAndServe(":8080", nil)
}
