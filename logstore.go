package gocollectanalytics

import (
	"encoding/json"
	"log"
)

// A Logstore is a basic Datastore that simply records to log,
// useful for debugging
type Logstore struct{}

// NewLogstore returns a pointer to a new Logstore
func NewLogstore() *Logstore {
	b := Logstore{}
	return &b
}

// LogIt serialises the supplied data to json and writes it to log
func (b *Logstore) LogIt(datatype interface{}) error {
	json, err := json.Marshal(datatype)
	if err != nil {
		return err
	}
	log.Printf("Logging: %+s", json)
	return nil
}
