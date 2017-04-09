package influxdb

import (
	"fmt"
	"github.com/influxdata/influxdb/client/v2"
	"github.com/lucasvasconcelos/criple-spider/config"
	"log"
	"strings"
	"time"
)

const (
	MyDB = "criple-spider"
)

var AppConfig *config.Spec = config.GetConfig()

func init() {
	if AppConfig.MetricsEnable {
		switch AppConfig.MetricsBackend {
		case "influxdb":
			if !strings.HasPrefix(AppConfig.MetricsHost, "http://") {
				config.PrintUsage()
				log.Fatal("MetricsHost must start with http://")
			}
		default:
			config.PrintUsage()
			log.Fatal("Need to set a MetricsBackend")
		}
	}

}

func SendMetric() {

	if !AppConfig.MetricsEnable {
		return
	}

	metricsHost := fmt.Sprintf("%v:%v", AppConfig.MetricsHost, AppConfig.MetricsPort)
	// Create a new HTTPClient
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Timeout: time.Second * 1,
		Addr:    metricsHost,
	})

	if err != nil {
		log.Println(err)
	}

	defer c.Close()

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "us",
	})
	if err != nil {
		log.Println(err)
	}

	// Create a point and add to batch
	tags := map[string]string{"pages": "visits"}
	fields := map[string]interface{}{
		"count": 1,
	}

	pt, err := client.NewPoint("pages", tags, fields, time.Now())
	if err != nil {
		log.Println(err)
	}
	bp.AddPoint(pt)

	// Write the batch
	if err := c.Write(bp); err != nil {
		log.Println(err)
	}

}
