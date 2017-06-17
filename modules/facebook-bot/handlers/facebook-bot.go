package handlers

import (
	"net/http"
	"scanbu-api/modules/search/lib"

	"fmt"

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
	fbRequest, _ := messenger.DecodeRequest(r)

	for _, entry := range fbRequest.Entry {
		for _, msg := range entry.Messaging {
			userID := msg.Sender.ID

			switch {
			case msg.Message != nil:
				search := msg.Message.Text

				products, err := lib.Search(search)
				if err != nil {
					msng.SendTextMessage(userID, "Tivemos um problema na busca do seu produto :(")
					msng.SendTextMessage(userID, "Tente novamente mais tarde :)")
				}
				gm := msng.NewGenericMessage(userID)
				for i, product := range products {
					if i == 2 || i == len(products)-1 {
						btn1 := msng.NewWebURLButton("Ver Mais", fmt.Sprintf("http://scanbu.com/search?keyword=%s", search))
						gm.AddNewElement(product.Message, "", product.Link, product.FullPicture, []messenger.Button{btn1})
						fmt.Println(product.Link)
						break
					}

					gm.AddNewElement(product.Message, "", product.Link, product.FullPicture, nil)
				}

				if len(products) == 0 {
					msng.SendTextMessage(userID, "Não encontramos nada :(")
				}

				msng.SendMessage(gm)
				return
			case msg.Delivery != nil:
				log.Println("Delivery received with content:", msg.Delivery)
			case msg.Postback != nil:
				log.Println("Postback received with content:", msg.Postback.Payload)
			}
		}
	}
}
