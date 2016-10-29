package main

/*
app.post('/webhook', (req, res) => {
	const data = req.body
	if (data.object === 'page') {
		data.entry.forEach(pageEntry => {
			pageEntry.messaging.forEach(messagingEvent => {
				if (messagingEvent.message) {
					if (!messagingEvent.message.is_echo) {
						handleMessage(messagingEvent)
					}
				}
			})
		})
		res.sendStatus(200)
	}
})
*/
import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func route(w http.ResponseWriter, req *http.Request) {

	fmt.Printf("%s", req.Body)
	w.WriteHeader(200)
	return
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
