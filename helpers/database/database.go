package database

import (
	"scanbu-api/helpers"

	log "github.com/Sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"
)

func init() {
	dbName := getDBName()
	setDB(dbName)
}

var databaseName string

// Session is the database session
var Session *mgo.Session

// StartDB start a database connection and return a session
func StartDB() *mgo.Session {
	strcon := getStrCon()

	var err error

	Session, err = mgo.Dial(strcon)
	if err != nil {
		panic(err)
	}

	return Session
}

// CloseDB close the database connection
func CloseDB() {
	Session.Close()
}

func getStrCon() (strcon string) {
	return helpers.GetENVorDefault("MONGODB_URI", "mongodb://localhost")
}

// DB return a mongodb database connection
func DB(dbName string) {
	Session.DB(dbName)
}

func getDBName() (strcon string) {
	return helpers.GetENVorDefault("DB_NAME", "scanbu-api")
}

func setDB(dbName string) {
	databaseName = dbName
}

// Collection return a mongodb collection
func Collection(collectionName string, indexes []mgo.Index) (c *mgo.Collection) {
	c = Session.DB(databaseName).C(collectionName)

	for _, index := range indexes {
		if err := c.EnsureIndex(index); err != nil {
			log.Fatal(err)
		}
	}

	return
}
