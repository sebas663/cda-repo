package db

import (
	"encoding/json"

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

// FindByQuery returns a single CdaContainer from the database.
func FindByQuery(typeDoc uint8, queryJSON string) ([]m.CdaContainer, error) {
	res := []m.CdaContainer{}

	var queryBson bson.M

	json.Unmarshal([]byte(queryJSON), &queryBson)

	if err := collection(typeDoc).Find(queryBson).All(&res); err != nil {
		return nil, err
	}

	return res, nil
}

// Save inserts an CdaContainer to the database.
func Save(a m.CdaContainer, typeDoc uint8) error {
	return collection(typeDoc).Insert(a)
}
