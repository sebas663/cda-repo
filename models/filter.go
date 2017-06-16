package models

import "time"

type (
	// FilterDB represents the structure of our resource
	FilterDB struct {
		DocumentID    string    `json:"documentID"`
		DocumentName  string    `json:"documentName"`
		EpisodeNumber uint32    `json:"episodeNumber"`
		EpisodeType   string    `json:"episodeType"`
		ReportDate    time.Time `json:"reportDate"`
		Nhc           uint32    `json:"nhc"`
		DocumentType  uint8     `json:"documentType"`
	}
	//FilterDBs array.
	FilterDBs []FilterDB
)
