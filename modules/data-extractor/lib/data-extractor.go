package lib

import (
	"fmt"
	"os"
	"os/signal"
	"scanbu-api/configs"
	"scanbu-api/modules/product/models"
	"syscall"

	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
	fb "github.com/huandu/facebook"
)

var wg sync.WaitGroup
var stop chan bool

func init() {
	stop = make(chan bool, 1)
	fb.Version = "v2.3"
}

func getGroupFeedByIDAndSince(groupID string, since time.Time) (feed []models.Product, err error) {
	// sinceStr := since.Format("2006-01-02")
	// sinceStr := "2017-01-01"

	// feed.until(2017 - 05 - 20).since(2017 - 05 - 19)

	groupPath := fmt.Sprintf("/%s", groupID)
	fieldsQuery := "feed{message,type,picture,full_picture,created_time,description,from,target,attachments{media},permalink_url}"

	res, err := fb.Get(groupPath, fb.Params{
		"fields":       fieldsQuery,
		"access_token": configs.FBToken,
	})

	if err == nil {
		res.DecodeField("feed.data", &feed)
	}

	return
}

func saveData(groupID string, since time.Time) {
	defer wg.Done()

	feeds, err := getGroupFeedByIDAndSince(groupID, since)
	if err != nil {
		log.Warn(err)
		return
	}

	for _, feed := range feeds {
		err = models.Products().Insert(feed)
		if err != nil {
			log.Warn(err)
			return
		}
	}
}

// ExtractorProcess is the data extractor process
func ExtractorProcess() {
	defer os.Exit(0)

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	ticker := time.NewTicker(time.Minute * 5)

	wg.Add(1)

	go func() {
		for range ticker.C {
			groupIDs := []string{
				"193939064109587",
				"1088976661131866",
				"415451778499368",
				"581149701917314",
				"331251590300019",
				"770134713017726",
				"619469771472103",
				"947261968650219",
				"959372877469093",
				"1450383971928186",
				"1475834732708100",
				"267063743450545",
				"1694177107530738",
				"357906320996128",
			}

			for _, groupID := range groupIDs {
				select {
				case <-stop:
					return
				default:
					wg.Add(1)
					go saveData(groupID, time.Now())
				}
			}
		}
	}()

	go func() {
		defer wg.Done()

		<-sigs
		stop <- true
		log.Info("System turning off")
		ticker.Stop()
	}()

	wg.Wait()
}
