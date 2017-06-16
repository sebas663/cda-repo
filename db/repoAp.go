package db

import (
	"time"

	m "../models"
	"gopkg.in/mgo.v2/bson"
)

// GetAll returns all CdaContainers from the database.
func GetAll(typeDoc uint8) ([]m.CdaContainer, error) {
	res := []m.CdaContainer{}

	if err := collection(typeDoc).Find(nil).All(&res); err != nil {
		return nil, err
	}

	return res, nil
}

// FindByDocumentID returns a single CdaContainer from the database.
func FindByDocumentID(mapa map[string]interface{}) ([]m.CdaContainer, error) {
	res := []m.CdaContainer{}
	documentID := mapa["documentID"]
	typeDoc := mapa["typeDoc"]
	query := mapa["query"]
	/*query := bson.M{
		"documentID":  documentID,
		"episodeType": bson.M{"$ne": "H"},
	}*/
	if err := collection(typeDoc).Find(query).All(&res); err != nil {
		return nil, err
	}

	return res, nil
}

// FindByDocumentName returns a single CdaContainer from the database.
func FindByDocumentName(id string) (*m.CdaContainer, error) {
	res := m.CdaContainer{}

	if err := collection(typeDoc).FindId(bson.ObjectIdHex(id)).One(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

// FindByEpisodeNumber returns a single CdaContainer from the database.
func FindByEpisodeNumber(id uint32) (*m.CdaContainer, error) {
	res := m.CdaContainer{}

	if err := collection(typeDoc).FindId(bson.ObjectIdHex(id)).One(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

// FindByReportDate returns a single CdaContainer from the database.
func FindByReportDate(from time.Time, to time.Time) (*m.CdaContainer, error) {
	res := m.CdaContainer{}

	if err := collection(typeDoc).FindId(bson.ObjectIdHex(id)).One(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

// FindByNhc returns a single CdaContainer from the database.
func FindByNhc(nhc uint32) (*m.CdaContainer, error) {
	res := m.CdaContainer{}

	if err := collection(typeDoc).FindId(bson.ObjectIdHex(id)).One(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Save inserts an CdaContainer to the database.
func Save(a m.CdaContainer) error {
	return collection(typeDoc).Insert(a)
}
