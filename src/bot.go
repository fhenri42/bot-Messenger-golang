
package main

import (
	"log"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"io/ioutil"
	"encoding/json"
)


type RecastRep struct {
	Message string `json:"message"`
	Results struct {
		Action struct {
			Done  bool   `json:"done"`
			Reply string `json:"reply"`
			Slug  string `json:"slug"`
		} `json:"action"`
		ConversationToken string `json:"conversation_token"`
		Entities          struct {
		} `json:"entities"`
		Intents []struct {
			Confidence float64 `json:"confidence"`
			Slug       string  `json:"slug"`
		} `json:"intents"`
		Language string `json:"language"`
		Memory   struct {
			Date       interface{} `json:"date"`
			RoomNumber interface{} `json:"room-number"`
		} `json:"memory"`
		NextActions []struct {
			Done  bool   `json:"done"`
			Reply string `json:"reply"`
			Slug  string `json:"slug"`
		} `json:"next_actions"`
		Replies   []string  `json:"replies"`
		Source    string    `json:"source"`
		Status    int64     `json:"status"`
		Uuid      string    `json:"uuid"`
		Version   string    `json:"version"`
	} `json:"results"`
}

func call_recast(msg string) string  {
	fmt.Printf("start call to recast")
	client := &http.Client{}

	form := url.Values{}
	form.Add("text", msg)
	req, err := http.NewRequest("POST","https://api.recast.ai/v2/converse" ,strings.NewReader(form.Encode()))
	if err != nil {
		log.Println(err)
		return "err"
	}
	req.Header.Set("Authorization", fmt.Sprintf("Token ADD_YOUR_TOKEN"))
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return "err"
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "err"
	}
	log.Println("\n \n ",string(body),"\n \n")
	var rep RecastRep
	err = json.Unmarshal(body, &rep)

	if err != nil {
		log.Println(err)
		return "err"
	}
	return rep.Results.Action.Reply

}

func  message_handler(data Data) {
	message := data.Entry[0].Messaging[0].Message.Text
	recipient := data.Entry[0].Messaging[0].Sender.ID
	msg := call_recast(message)
	post_facebook(msg, recipient)
}
