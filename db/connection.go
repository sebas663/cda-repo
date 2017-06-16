package db

import (
	"log"

	"gopkg.in/mgo.v2"
)

var db *mgo.Database

func init() {

	session, err := mgo.Dial("127.0.0.1/cda_repo")

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db = session.DB("cda_repo")
}

func epicrisis() *mgo.Collection {
	return db.C("epicrisis")
}

func aps() *mgo.Collection {
	return db.C("aps")
}

func imagenes() *mgo.Collection {
	return db.C("imagenes")
}

func laboratorios() *mgo.Collection {
	return db.C("laboratorios")
}

func collection(typeDoc uint8) *mgo.Collection {
	if typeDoc == 1 {
		return epicrisis()
	}
	if typeDoc == 8 {
		return aps()
	}
	if typeDoc == 13 {
		return imagenes()
	}
	if typeDoc == 11 {
		return laboratorios()
	}
}
