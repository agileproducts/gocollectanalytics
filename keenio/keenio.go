package keenio

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

// NewKeenIOStore returns a pointer to a Keenio store based on supplied config
func NewKeenIOStore(kc KeenIOConfig) (*KeenIOStore, error) {
	if kc.WriteKey == "" {
		return nil, errors.New("Cannot use KeenIO store, write key not supplied")
	}
	if kc.ProjectID == "" {
		return nil, errors.New("Cannot use KeenIO store, project ID not supplied")
	}
	return &KeenIOStore{config: kc}, nil
}

// logs supplied datapoint data to keenio
func (ks *KeenIOStore) LogIt(datatype interface{}) error {

	//serialize
	json, err := json.Marshal(datatype)

	if err != nil {
		return err
	}

	//for the moment the use of a single collection called 'Events' is hardcoded in
	keensays, err := ks.send(json, "/events/Events")

	if err != nil {
		return err
	}

	message, err := ioutil.ReadAll(keensays.Body)
	if err != nil {
		return err
	}

	log.Print(keensays.StatusCode)
	log.Printf("%+s", message)
	return nil

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
