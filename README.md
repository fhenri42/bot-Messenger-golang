# README is not up to date (messing the config file part :( )!
# Start coding your bot: Recast.AI + Facebook messenger

* Here, you'll learn how to build a bot with Facebook Messenger and Recast.AI.

## Requirements
* Create an account on [Recast.AI](https://recast.ai/signup).
* Create an account on [Facebook Developers](https://developers.facebook.com/). (same account than your personal Facebook account)

## Set up your Recast.AI account

##### Create your bot

* Log into your [Recast.AI](https://recast.ai/login) account.
* Create a new bot.

##### Get your token

* In your profile, click your bot.
* In the tab-menu, click on the the little screw.
* Here is the `request access token` you will need to configure your bot.

## Create your facebook page
* [Create your page](https://www.facebook.com/pages/create/?ref_type=logout_gear).

 [facebook]: https://raw.githubusercontent.com/RecastAI/bot-messenger/master/ressources/S%C3%A9lection_021.png "Creating you page"

![alt text][facebook]
* Choose the category of your page.
* Fill out the Facebook requirements step by step.

[facebook-set-up]: https://raw.githubusercontent.com/RecastAI/bot-messenger/master/ressources/S%C3%A9lection_022.png "Steup of your page"

![alt text][facebook-set-up]

## Set up your facebook account

* Log on to your Facebook Developers account.
* Create a new Facebook app.

[facebook-first]: https://raw.githubusercontent.com/RecastAI/bot-messenger/master/ressources/S%C3%A9lection_028.png "first page"
![alt text][facebook-first]


* Get your app secret and ID [Dashboard](https://developers.facebook.com/apps/258158857911674/dashboard/).

[facebook-app]: https://raw.githubusercontent.com/RecastAI/bot-messenger/master/ressources/S%C3%A9lection_025.png  "Creating you page"

![alt text][facebook-app]

* Get your page Token [Messanger](https://developers.facebook.com/apps/258158857911674/messenger/).

[facebook-pageToken]: https://raw.githubusercontent.com/RecastAI/bot-messenger/master/ressources/S%C3%A9lection_026.png "Creating you page"

![alt text][facebook-pageToken]

## Start your bot in local
```bash
git clone https://github.com/fhenri42/bot-Messenger-golang.git
```

#### Ngrok

* Download the appropriate version of [Ngrok](https://ngrok.com/download).
* Open a new tab in your terminal:
```
./ngrok http 5000
```
* Copy past the ``` https://*******ngrok.io``` you get, you will need it for the next step.
* Leave your Ngrok serveur running.

## Complete the config.go

* Copy your Recast.AI `Recast.AI access token`
* Copy your page access Token `Token of your Page`
* Copy your validationToken `The token of your Webhook`

* Incoming :)

## Launching your Bot

* make sure to have ngrok launched and the correct URL in you config file.

```bash
 go run src/*.go
```
#### Config webhook

* go back to the Facebook Developer page and add a new webhook.

[webhook]: https://blog.recast.ai/wp-content/uploads/2016/09/S%C3%A9lection_020.png "Webhook page"

![alt text][webhook]
* Subscribe your page to the webhook you just created.

[suscribe]: https://raw.githubusercontent.com/RecastAI/bot-messenger/master/ressources/S%C3%A9lection_024.png "Subscribe page"

![alt text][suscribe]

## Result

[result]: https://raw.githubusercontent.com/RecastAI/bot-messenger/master/ressources/S%C3%A9lection_023.png

![alt text][result]

## Your bot
```go

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
```
* Have fun coding your bot! :)

## Author

Henri Floren - Recast.AI
henri.floren@recast.ai
