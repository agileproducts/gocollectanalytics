package gocollectanalytics

import (
	"testing"
)

func TestBogstore(t *testing.T) {
	bog := NewBogstore()
	err := bog.LogIt("moo")
	if err != nil {
		t.Fatal("Problem writing to Log")
	}
}

func TestKollector(t *testing.T) {
	st := NewBogstore()
	koll := NewKollector(st)
	event := testEvent()
	message := koll.record(event)
	if message != "ok" {
		t.Fatal("expected", "ok", "got", message)
	}
}

// https://blog.codeship.com/testing-in-go/
