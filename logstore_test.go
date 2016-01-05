package gocollectanalytics

import (
	"testing"
)

func TestLogstore(t *testing.T) {
	ls := NewLogstore()
	event := testEvent()
	err := ls.LogIt(event)
	if err != nil {
		t.Fatal("Problem writing to Log")
	}
}
