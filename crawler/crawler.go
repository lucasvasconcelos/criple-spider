package crawler

import (
	"github.com/lucasvasconcelos/criple-spider/config"
	"github.com/lucasvasconcelos/criple-spider/database/couchdb"
	"github.com/lucasvasconcelos/criple-spider/metrics/influxdb"
	"log"
	"time"
)

var AppConfig *config.Spec = config.GetConfig()

// Throttle the http.Get
func throttleCrawl(max int) (sem chan int) {
	sem = make(chan int, max)
	for i := 0; i < max; i++ {
		sem <- 1
	}
	return sem
}

func StartCrawling(startUrl string) {
	toVisit := make(chan string)

	go func() { toVisit <- startUrl }()

	throttle := throttleCrawl(AppConfig.MaxConcurrency)

	for url := range toVisit {
		go func() {
			<-throttle
			visitPage(url, toVisit)
			throttle <- 1
		}()
	}
}

func visitPage(url string, toVisit chan string) {
	time.Sleep(500 * time.Millisecond)
	visitedPage, err := couchdb.WasVisited(url)
	if err != nil {
		err := visitedPage.ParseLinks()

		if err != nil {
			log.Println(err)
		}

		err = couchdb.Save(&visitedPage)

		if err == nil {
			log.Printf("%v was visited\n", visitedPage.Url)
			influxdb.SendMetric()
		}
	}

	if len(visitedPage.Links) != 0 {
		go func() {
			for _, link := range visitedPage.Links {
				_, err := couchdb.WasVisited(link)
				if err != nil || link != AppConfig.StartUrl {
					toVisit <- link
				}
			}
		}()
	}
}
