package handlers

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	messenger "github.com/mileusna/facebook-messenger"
)

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
	fbRequest, err := messenger.DecodeRequest(r)
	if err != nil {
		return
	}

	for _, entry := range fbRequest.Entry {
		for _, msg := range entry.Messaging {
			userID := msg.Sender.ID

			switch {
			case msg.Message != nil:
				log.Info("Msg received with content:", msg.Message.Text)
				msng.SendTextMessage(userID, "Hello there")

				gm := msng.NewGenericMessage(userID)
				gm.AddNewElement("Title", "Subtitle", "http://mysite.com", "http://mysite.com/some-photo.jpeg", nil)

				btn1 := msng.NewWebURLButton("Contact US", "http://mysite.com/contact")
				btn2 := msng.NewPostbackButton("Ok", "THIS_DATA_YOU_WILL_RECEIVE_AS_POSTBACK_WHEN_USER_CLICK_THE_BUTTON")
				gm.AddNewElement("Site title", "Subtitle", "http://mysite.com", "http://mysite.com/some-photo.jpeg", []messenger.Button{btn1, btn2})

				// ok, message is ready, lets send
				msng.SendMessage(gm)
		}
	}
}
