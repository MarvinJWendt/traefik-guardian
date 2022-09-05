package db

import (
	"github.com/asdine/storm/v3"
	"github.com/sirupsen/logrus"
)

var db *storm.DB

func Setup() error {
	dbTmp, err := storm.Open("/data/credentials.db")
	db = dbTmp
	return err
}

func Close() {
	db.Close()
}

func GetDB() *storm.DB {
	if db == nil {
		logrus.Fatal("db not initialized")
	}
	return db
}
