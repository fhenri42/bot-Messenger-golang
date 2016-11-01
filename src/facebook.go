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
	rep,body,err := request.Post("https://graph.facebook.com/v2.6/me/messages?token=ADD_YOUR_TOKEN").Send(send).End()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("\nrep = ", rep)
	log.Println("\nbody = ", body)

}
