package gocollectanalytics

import (
	"log"
	"os"
	"testing"
)

func TestItRequiresAWriteKey(t *testing.T) {
	kc := KeenIOConfig{WriteKey: "", ProjectID: "1243"}
	_, err := newKeenIOStore(kc)
	if err == nil {
		t.Fatalf("It should require a write key when initializing a keenIO store")
	}
}

func TestItRequiresAProjectID(t *testing.T) {
	kc := KeenIOConfig{WriteKey: "1232", ProjectID: ""}
	_, err := newKeenIOStore(kc)
	if err == nil {
		t.Fatalf("It should require a project ID when initializing a keenIO store")
	}
}

/*
func TestItSavesAnEvent(t *testing.T) {
	ks := keenIntegrationTestStore()
	event := testEvent()
	ks.logEvent(event)
}
*/

func keenIntegrationTestStore() *KeenIOStore {
	kc := KeenIOConfig{WriteKey: os.Getenv("KEENIO_WRITE_KEY"), ProjectID: os.Getenv("KEENIO_PROJECT_ID")}
	ks, err := newKeenIOStore(kc)
	if err != nil {
		log.Fatalf("Error initializing integration test store. Missing environment variables?")
	}
	return ks
}

func testEvent() Event {
	return Event{Site: "BuildTest", Category: "TestCategory", Action: "TestAction"}
}
