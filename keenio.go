package gocollectanalytics

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

const (
	baseURL = "https://api.keen.io/3.0/projects/"
)

// KeenIOConfig holds connection data for Keen IO
type KeenIOConfig struct {
	WriteKey  string
	ProjectID string
}

// A KeenIOStore records to Keen IO
type KeenIOStore struct {
	config     KeenIOConfig
	httpClient http.Client
}

//KeenIOCollector implements an http handler function for receiving data
//and sending it to the project specified in the supplied config
func KeenIOCollector(kc KeenIOConfig) (collector *Collector, err error) {

	ks, err := newKeenIOStore(kc)

	if err != nil {
		return
	}

	collector = new(Collector)
	collector.Store = ks
	return
}

func newKeenIOStore(kc KeenIOConfig) (*KeenIOStore, error) {
	if kc.WriteKey == "" {
		return nil, errors.New("Cannot use KeenIO store, write key not supplied")
	}
	if kc.ProjectID == "" {
		return nil, errors.New("Cannot use KeenIO store, project ID not supplied")
	}
	return &KeenIOStore{config: kc}, nil
}

func (ks *KeenIOStore) logDatapoint(datatype interface{}) {

	json, err := json.Marshal(datatype)

	if err != nil {
		log.Print(err)
	}

	ks.send(json, "/events")
}

func (ks *KeenIOStore) send(json []byte, path string) {

	// construct url
	url := baseURL + ks.config.ProjectID + path
	log.Printf("%+s", url)
}
