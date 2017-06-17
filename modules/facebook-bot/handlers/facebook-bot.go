package handlers

import (
	"log"
	"net/http"

	messenger "github.com/mileusna/facebook-messenger"
)

// use public messenger for simpler code demonstration
var msng = &messenger.Messenger{
	AccessToken: "EAAS1E2CpgCYBAFa5FqXcgsmCpbDfjrnDLp3EUbPQjTINi2Dae7CSKqYeaZBloGx3ZAnZAOVU2W4RrtCA9oUIS4xCzzCRKZBsOdlz5KtCDKJGhEZAxwmu7gZCqHZAwRwI7bXJfEl4z4lk9q55nzE8ZCC87z6saML24VBc7CZCGwQi4TzCYJrIzXuuH",
	PageID:      "1871739449816952",
}

// FacebookBot is the facebook bot webhook handler
func FacebookBot(w http.ResponseWriter, r *http.Request) {
	msng.VerifyWebhook(w, r)
}

// FacebookBotReceiver is the facebook bot message receiver handler
func FacebookBotReceiver(w http.ResponseWriter, r *http.Request) {
	fbRequest, _ := messenger.DecodeRequest(r) // decode entire request received from Facebook into FacebookRequest struct

	// now you have it all and you can do whatever you want with received request
	// enumerate each entry, and each message in entry
	for _, entry := range fbRequest.Entry {
		// pageID := entry.ID  // here you can find page id that received message
		for _, msg := range entry.Messaging {
			userID := msg.Sender.ID // user that sent you a message

			// but "message" can be text message, delivery report or postback, so check it what it is
			// it can only be one of this, so we use switch
			switch {
			case msg.Message != nil:
				log.Println("Msg received with content:", msg.Message.Text)
				msng.SendTextMessage(userID, "Hello there")
				// check First example for more sending messages examples

			case msg.Delivery != nil:
				// delivery report received, check First example what to do next

			case msg.Postback != nil:
				// postback received, check First example what can you do with that
				log.Println("Postback received with content:", msg.Postback.Payload)
			}
		}
	}
}
