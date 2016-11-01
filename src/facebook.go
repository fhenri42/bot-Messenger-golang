package main

import (
	"log"
	"github.com/parnurzeal/gorequest"
)

type FacebookRes struct {

	Message struct {
		Text string `json:"text"`
	} `json:"message"`
	Recipient struct {
		ID int64 `json:"id"`
	} `json:"recipient"`

}
func post_facebook(msg string, id int64 ) {

	request := gorequest.New()
	var send FacebookRes
	send.Message.Text = msg
	send.Recipient.ID = id
	rep,body,err := request.Post("https://graph.facebook.com/v2.6/me/messages?access_token=EAADqy0sFfXoBAKVdZCWPo0b8JMDD8pHHO8i3iRA7FQQkhkmVaGyTgBdQxUHqHzfkWYbPXRTOSAklZBZAj7oU3ZAL3sE0Lm1O8WZADiiXLMjFqd84p8NR03ACmbnGICa1ydm5eHV7ZCXkozmhbv4ZCKZAeRKHvvhbDaTqYUtETY6eWgZDZD").Send(send).End()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("rep = ", rep)
	log.Println("body = ", body)

}
