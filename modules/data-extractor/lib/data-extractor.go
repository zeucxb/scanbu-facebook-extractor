package lib

import (
	"fmt"
	"scanbu-api/configs"
	"scanbu-api/modules/product/models"

	"sync"
	"time"

	fb "github.com/huandu/facebook"
)

var wg sync.WaitGroup

func getGroupFeedByIDAndSince(groupID string, since time.Time) (feed []models.Product, err error) {
	sinceStr := since.Format("2006-01-02")

	// feed.until(2017 - 05 - 20).since(2017 - 05 - 19)

	groupPath := fmt.Sprintf("/%s", groupID)
	fieldsQuery := fmt.Sprintf("feed.since(%s){message,type,picture,full_picture,created_time,description,from,target,attachments{media},permalink_url}", sinceStr)

	res, err := fb.Get(groupPath, fb.Params{
		"fields":       fieldsQuery,
		"access_token": configs.FBToken,
	})

	if err == nil {
		res.DecodeField("feed.data", &feed)
	}

	return
}

func saveData(groupID string, since time.Time, async bool) {
	if async {
		defer wg.Done()
	}

	feeds, err := getGroupFeedByIDAndSince(groupID, since)
	if err != nil {
		return
	}

	for _, feed := range feeds {
		err = models.Products().Insert(feed)

		if err == nil {
			fmt.Println("ok")
			fmt.Println(feed.Attachments)
		}
		fmt.Println(err)
	}
}

func Proccess(groupIDs []string) {
	for _, groupID := range groupIDs {
		wg.Add(1)
		go saveData(groupID, time.Now(), true)
	}

	wg.Wait()
}
