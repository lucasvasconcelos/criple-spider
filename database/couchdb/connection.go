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
var AppConfig *config.Spec = config.GetConfig()

const (
	MyDB = "criple-spider"
)

func init() {
	var timeout = time.Duration(500 * time.Millisecond)
	conn, _ := couchdb.NewConnection(AppConfig.DbHost, AppConfig.DbPort, timeout)
	auth := couchdb.BasicAuth{Username: AppConfig.DbUser, Password: AppConfig.DbPassword}

	dbList, err := conn.GetDBList()

	if err != nil {
		config.PrintUsage()
		log.Fatal(err)
	}

	if !dbExists(dbList) {
		conn.CreateDB(MyDB, &auth)
	}

	db = conn.SelectDB(MyDB, &auth)
}

func dbExists(dbList []string) bool {
	for _, db := range dbList {
		if db == MyDB {
			return true
		}
	}
	return false
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
