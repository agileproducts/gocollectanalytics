package keenio

import (
	_ "os"
	"testing"
)

func TestItRequiresAWriteKey(t *testing.T) {
	kc := KeenIOConfig{WriteKey: "", ProjectID: "1243"}
	_, err := NewKeenIOStore(kc)
	if err == nil {
		t.Fatalf("It should require a write key when initializing a keenIO store")
	}
}

func TestItRequiresAProjectID(t *testing.T) {
	kc := KeenIOConfig{WriteKey: "1232", ProjectID: ""}
	_, err := NewKeenIOStore(kc)
	if err == nil {
		t.Fatalf("It should require a project ID when initializing a keenIO store")
	}
}

/*func TestKeenIOStore(t *testing.T) {
	kc := KeenIOConfig{WriteKey: os.Getenv("KEENIO_WRITE_KEY"), ProjectID: os.Getenv("KEENIO_PROJECT_ID")}
	ks, _ := NewKeenIOStore(kc)
	event := testEvent()
	err := ks.LogIt(event)
	if err != nil {
		t.Fatal("Problem")
	}
}*/

/*func TestItSendsDataToKeenIO(t *testing.T) {
	ks := keenIntegrationTestStore(t)
	event := testEvent()
	ks.logDatapoint(event)
}

func keenIntegrationTestStore(t *testing.T) *KeenIOStore {
	kc := KeenIOConfig{WriteKey: os.Getenv("KEENIO_WRITE_KEY"), ProjectID: os.Getenv("KEENIO_PROJECT_ID")}
	ks, err := newKeenIOStore(kc)
	if err != nil {
		t.Fatal("Error initializing integration test store. Missing environment variables?")
	}
	return ks
} */
