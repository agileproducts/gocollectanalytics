package gocollectanalytics

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
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

// KeenIOCollector implements an http handler function for receiving data
// and sending it to the project specified in the supplied config
func KeenIOCollector(kc KeenIOConfig) (collector *Collector, err error) {

	ks, err := newKeenIOStore(kc)

	if err != nil {
		return
	}

	collector = new(Collector)
	collector.Store = ks
	return
}

// sets up Keenio store based on supplied config
func newKeenIOStore(kc KeenIOConfig) (*KeenIOStore, error) {
	if kc.WriteKey == "" {
		return nil, errors.New("Cannot use KeenIO store, write key not supplied")
	}
	if kc.ProjectID == "" {
		return nil, errors.New("Cannot use KeenIO store, project ID not supplied")
	}
	return &KeenIOStore{config: kc}, nil
}

// logs supplied datapoint data to keenio
func (ks *KeenIOStore) logDatapoint(datatype interface{}) {

	//serialize
	json, err := json.Marshal(datatype)

	log.Printf("%+s", json)

	if err != nil {
		log.Print(err)
	}

	//for the moment the use of a single collection called 'Events' is hardcoded in
	keensays, err := ks.send(json, "/events/Events")

	if err != nil {
		log.Print(err)
	}

	message, err := ioutil.ReadAll(keensays.Body)
	if err != nil {
		log.Print(err)
	}

	log.Print(keensays.StatusCode)
	log.Printf("%+s", message)

}

func (ks *KeenIOStore) send(json []byte, path string) (*http.Response, error) {

	// construct url
	url := baseURL + ks.config.ProjectID + path

	// assemble http request
	req, err := http.NewRequest("POST", url, bytes.NewReader(json))
	if err != nil {
		return nil, err
	}

	// add auth
	req.Header.Add("Authorization", ks.config.WriteKey)

	// set length/content-type
	if json != nil {
		req.Header.Add("Content-Type", "application/json")
		req.ContentLength = int64(len(json))
	}

	return ks.httpClient.Do(req)

}
