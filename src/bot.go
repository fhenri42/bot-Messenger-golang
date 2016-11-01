
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
}

func call_recast(msg string) {
	fmt.Printf("start call to recast")
	client := &http.Client{}

	form := url.Values{}
	form.Add("text", msg)
	req, err := http.NewRequest("POST","https://api.recast.ai/v2/converse" ,strings.NewReader(form.Encode()))
	if err != nil {
		fmt.Printf("in the err\n")
		log.Println(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Token 8eb71c44150033815807a532db822e59"))
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("in the err\n")
		log.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("in the err\n")
		log.Println(err)
	}
	var rep RecastRep
	err = json.Unmarshal(body, &rep)

	if err != nil {
		fmt.Printf("in the err\n")
		log.Println(err)
	}
	log.Println("\n",rep)


}

func  message_handler(data Data) {
	message := data.Entry[0].Messaging[0].Message.Text
	recipient := data.Entry[0].Messaging[0].Recipient.ID
	fmt.Printf("message = %s\n",message)
	fmt.Printf("recipient = %d\n", recipient)
	call_recast(message)
}
