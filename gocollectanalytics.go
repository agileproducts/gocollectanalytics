/*
Package gocollectanalytics implements an http handler function to collect data according to
Google Analytics -style styntax and save it in the desired datastore
*/
package gocollectanalytics

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
)

// A Collector provides handling for receiving data and recording it to
// the given store
type Collector struct {
	Store Datastore
}

// A Datastore is any place to store data. It must satisfy this interface,
// with a logDatapoint() method that accepts a data type
type Datastore interface {
	logDatapoint(interface{})
}

// A LogStore is a datastore that simply records to log, useful for debugging
type LogStore struct{}

//LogCollector returns a collector that just writes to a logfile datastore
func LogCollector() (collector *Collector, err error) {

	err = nil

	collector = new(Collector)
	collector.Store = new(LogStore)

	return

}

// CollectData is a http.HandlerFunc to parse and validate querystring data
// then save it as the appropriate type in the specified datastore.
func (s *Collector) CollectData(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	dataValid, err := validateParameters(params)
	if dataValid != true {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		e := createEvent(params)
		go s.Store.logDatapoint(e) // naive concurrency: http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang/
		w.WriteHeader(http.StatusOK)
		//fmt.Fprint(w, "collected ok")
	}
}

func (ls LogStore) logDatapoint(datatype interface{}) {

	json, err := json.Marshal(datatype)

	if err != nil {
		log.Print(err)
	}

	log.Printf("%+s", json)
}

// An Event is a user interactions with content that can be tracked independently
// from a web page or a screen load. A simple example would be clicking a link.
type Event struct {
	Site     string `json:"site"`
	ClientID string `json:"clientid"`
	Category string `json:"category"`
	Action   string `json:"action"`
	Label    string `json:"label"`
	Value    int    `json:"value"`
}

// mulitpleErrors are a slice of Errors
type multipleErrors []error

// createEvent turns the data paramaters associated with a hit type of 'event'
// into a golang type Event
func createEvent(data url.Values) Event {
	e := Event{Site: data.Get("tid"), Category: data.Get("ec"), Action: data.Get("ea"), Label: data.Get("el")}
	//skipping value and client id for now
	return e
}

// validateParameters takes a set of url.Values and parses them to ensure they
// match the required data specification.
func validateParameters(vals url.Values) (bool, multipleErrors) {

	errs := []error{}

	// v is the measurement protocol version. The only current value is '1'
	if vals.Get("v") != "1" {
		errs = append(errs, errors.New("Version v must equal 1"))
	}

	// tid is the site id of the client we are collecting data from
	if vals.Get("tid") == "" {
		errs = append(errs, errors.New("Site id tid must be supplied"))
	}

	// NOT IMPLEMENTED - cid is the browser id of the client we are collecting data from
	//if vals.Get("tid") == "" {
	//  errs = append(errs, errors.New("Client id tid must be supplied"))
	//}

	// t is the type of hit. At present the only supported values is 'event'
	if vals.Get("t") != "event" {
		errs = append(errs, errors.New("Hit type 't' must be set, only type 'event' is currently supported"))
	}

	// events must have a category ec
	if vals.Get("t") == "event" && vals.Get("ec") == "" {
		errs = append(errs, errors.New("Events must have a category"))
	}

	// events must have an action ea
	if vals.Get("t") == "event" && vals.Get("ea") == "" {
		errs = append(errs, errors.New("Events must have an action"))
	}

	if len(errs) > 0 {
		return false, errs
	}
	return true, nil

}

// collects Errors into multipleErrors
func (e multipleErrors) Error() string {
	if len(e) == 1 {
		return e[0].Error()
	}

	msg := "multiple errors:"
	for _, err := range e {
		msg += "\n" + err.Error()
	}
	return msg
}
