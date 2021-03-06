package models

import "gopkg.in/mgo.v2/bson"
import "time"

type (
	// CdaContainer represents the structure of our resource
	CdaContainer struct {
		ID            bson.ObjectId `json:"id" bson:"_id,omitempty"`
		DocumentID    string        `json:"documentID" bson:"documentID"`
		DocumentName  string        `json:"documentName" bson:"documentName"`
		EpisodeNumber uint32        `json:"episodeNumber" bson:"episodeNumber"`
		EpisodeType   string        `json:"episodeType" bson:"episodeType"`
		ReportDate    time.Time     `json:"reportDate" bson:"reportDate"`
		Nhc           uint32        `json:"nhc" bson:"nhc"`
		Cda           string        `json:"cda" bson:"cda"`
		DocumentType  uint8         `json:"documentType" bson:"documentType"`
	}
	//CdaContainers array.
	CdaContainers []CdaContainer
)
