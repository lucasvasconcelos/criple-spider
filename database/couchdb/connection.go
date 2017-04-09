package couchdb

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/lucasvasconcelos/criple-spider/config"
	"github.com/lucasvasconcelos/criple-spider/parse"
	"github.com/rhinoman/couchdb-go"
	"log"
	"time"
)

var db *couchdb.Database

func init() {
	var timeout = time.Duration(500 * time.Millisecond)
	AppConfig := config.GetConfig()
	conn, _ := couchdb.NewConnection(AppConfig.DbHost, AppConfig.DbPort, timeout)
	auth := couchdb.BasicAuth{Username: AppConfig.DbUser, Password: AppConfig.DbPassword}

	_, err := conn.GetDBList()

	if err != nil {
		config.PrintUsage()
		log.Fatal(err)
	}

	db = conn.SelectDB("criple-spider", &auth)
}

func Save(htmlpage *parse.HTMLPage) error {
	sha := sha256.Sum256([]byte(htmlpage.Url))
	id := hex.EncodeToString(sha[:])
	_, err := db.Save(htmlpage, id, "")
	if err != nil {
		return err
	}
	return nil
}

func WasVisited(uri string) (parse.HTMLPage, error) {
	var h parse.HTMLPage
	h.Url = uri
	sha := sha256.Sum256([]byte(h.Url))
	id := hex.EncodeToString(sha[:])
	_, err := db.Read(id, &h, nil)
	if err != nil {
		return h, err
	}
	return h, nil
}
