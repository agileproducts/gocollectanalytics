package gocollectanalytics

import (
	"github.com/agileproducts/gocollectanalytics"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
	"testing"
)

func TestItRequiresAWriteKey(t *testing.T) {
	_, err := newKeenIOStore("", "12345")
	if err == nil {
		t.Fatalf("It should require a write key when initializing a keenIO store")
	}
}

func TestItRequiresAProjectID(t *testing.T) {
	_, err := newKeenIOStore("12345", "")
	if err == nil {
		t.Fatalf("It should require a project ID when initializing a keenIO store")
	}
}

func TestItSavesAnEvent(t *testing.T) {
	ks := keenIntegrationTestStore()
	event := testEvent()
	ks.logEvent(event)
}

func keenIntegrationTestStore() *KeenIOStore {
	ks, err := newKeenIOStore(os.Getenv("KEENIO_WRITE_KEY"), os.Getenv("KEENIO_PROJECT_ID"))
	log.Printf("bob: %v", os.Getenv("KEENIO_WRITE_KEY"))
	if err != nil {
		log.Fatalf("Error initializing integration test store. Missing environment variables?")
	}
	return ks
}

func testEvent() gocollectanalytics.Event {
	return gocollectanalytics.Event{Site: "BuildTest", Category: "TestCategory", Action: "TestAction"}
}
