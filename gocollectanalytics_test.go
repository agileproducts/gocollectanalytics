package gocollectanalytics

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var validURL = "/collect?v=1&tid=test&t=event&ec=test&ea=test"
var validData, _ = url.ParseQuery(validURL)

func TestItDemandsAVersion(t *testing.T) {
	data := copyURLValues(validData)
	data.Del("v")
	checkItDoesntValidate(data, "It should say data without a v parameter is invalid", t)
}

func TestItDemandsASiteId(t *testing.T) {
	data := copyURLValues(validData)
	data.Del("tid")
	checkItDoesntValidate(data, "It should say data without a tid parameter is invalid", t)
}

func TestItDemandsAHitType(t *testing.T) {
	data := copyURLValues(validData)
	data.Del("t")
	checkItDoesntValidate(data, "It should say data without a t parameter is invalid", t)
}

func TestItOnlyAcceptsVersion1(t *testing.T) {
	data := copyURLValues(validData)
	data.Set("v", "2")
	checkItDoesntValidate(data, "It should say data where v != '1' is invalid", t)
}

//For now
func TestItOnlyAcceptsEvents(t *testing.T) {
	data := copyURLValues(validData)
	data.Set("t", "pageview")
	checkItDoesntValidate(data, "It should say data where t != 'event' is invalid", t)
}

func TestEventsMustHaveACategory(t *testing.T) {
	data := copyURLValues(validData)
	data.Del("ec")
	checkItDoesntValidate(data, "It should say event data without a category parameter is invalid", t)
}

func TestEventsMustHaveAnAction(t *testing.T) {
	data := copyURLValues(validData)
	data.Del("ea")
	checkItDoesntValidate(data, "It should say event data without an action parameter is invalid", t)
}

func TestItMakesAnEvent(t *testing.T) {
	data := copyURLValues(validData)
	event := createEvent(data)
	if event.Category != data.Get("ec") {
		t.Fatalf("It should be able to make an event from a set of valid URL params %v", data)
	}
	if event.Action != data.Get("ea") {
		t.Fatalf("It should be able to make an event from a set of valid URL params %v", data)
	}
}

/*func TestItSavesAnEventToLog(t *testing.T) {
	data := copyURLValues(validData)
	coll, err := NewCollector("log")
	if err != nil {
		t.Fatalf("It should be able to save an event to log")
	}
	event := createEvent(data)
	bob := coll.saveEvent(event)
	if bob != "log" { //rubbish test
		t.Fatalf("It should be able to save an event to log")
	}
}*/

func TestCollectData(t *testing.T) {
	coll, err := NewCollector("log")
	if err != nil {
		t.Fatalf("It should be able to collect data")
	}
	handler := coll.CollectData
	req, _ := http.NewRequest("GET", validURL, nil)
	w := httptest.NewRecorder()
	handler(w, req)

	checkItWasOK(w, t)
}

func checkItDoesntValidate(data url.Values, message string, t *testing.T) {
	isOk, _ := validateParameters(data)
	if isOk == true {
		t.Fatalf("%v", message)
	}
}

func copyURLValues(m map[string][]string) url.Values {
	newMap := map[string][]string{}
	for k, v := range m {
		newMap[k] = v
	}
	return newMap
}

func checkItWasOK(w *httptest.ResponseRecorder, t *testing.T) {
	if w.Code != http.StatusOK {
		t.Fatalf("Unexpected status, expected 200 but got %d", w.Code)
	}
}
