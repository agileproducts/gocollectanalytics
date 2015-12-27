package gocollectanalytics

import (
	"errors"
	"github.com/agileproducts/gocollectanalytics"
	"log"
	"net/http"
)

const (
	baseURL = "https://api.keen.io/3.0/projects/"
)

// A KeenIOStore records to Keen IO
type KeenIOStore struct {
	writeKey   string
	projectID  string
	httpClient http.Client
}

func newKeenIOStore(wk string, pid string) (*KeenIOStore, error) {
	if wk == "" {
		return nil, errors.New("Cannot use KeenIO store, write key not supplied")
	}
	if pid == "" {
		return nil, errors.New("Cannot use KeenIO store, project ID not supplied")
	}
	return &KeenIOStore{writeKey: wk, projectID: pid}, nil
}

func (ks *KeenIOStore) logEvent(e gocollectanalytics.Event) {
	log.Printf("Saving to %s %+v", ks, e)
}
