package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

// the specification of the app config
type Spec struct {
	StartUrl        string `default:"https://news.ycombinator.com" required:"false"`
	DbHost          string `default:"localhost" required:"false"`
	DbPort          int    `default:"5984" required:"false"`
	DbUser          string `default:"admin" required:"false"`
	DbPassword      string `default:"password" required:"false"`
	DbBackend       string `default:"couchdb" required:"false"`
	MetricsEnable   bool   `default:"false" required:"false"`
	MetricsHost     string `default:"http://localhost" required:"false"`
	MetricsPort     int    `default:"8086" required:"false"`
	MetricsUser     string `default:"admin" required:"false"`
	MetricsPassword string `default:"password" required:"false"`
	MetricsBackend  string `default:"influxdb" required:"false"`
	MaxConcurrency  int    `default:"2" required:"false"`
}

var AppConfig Spec

func init() {
	LoadConfig()
}

func LoadConfig() {
	err := envconfig.Process("criple_spider", &AppConfig)
	if err != nil {
		PrintUsage()
		log.Fatal(err)
	}
}

func GetConfig() *Spec {
	return &AppConfig
}

func PrintUsage() {
	envconfig.Usage("criple_spider", &AppConfig)
}
